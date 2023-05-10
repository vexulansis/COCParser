package pool

import (
	"context"
	"sync"
	"time"

	. "github.com/vexulansis/COCParser/pkg/client"
	. "github.com/vexulansis/COCParser/pkg/worker"
)

type Pool struct {
	Name    string
	Config  *PoolConfig
	Manager *PoolManager
	Workers []*Worker
	WG      *sync.WaitGroup
	Outputs []chan []byte
	Inputs  []chan []byte
}
type PoolConfig struct {
	Client  Client
	Process ProcessFunc
	Delay   int
}

// Creates Pool example
func NewPool(name string, config *PoolConfig) *Pool {
	pool := &Pool{
		Name:   name,
		Config: config,
		WG:     &sync.WaitGroup{},
	}
	pool.Manager = NewPoolManager(pool)
	return pool
}

// Creates workers and adds them to Pool
func (p *Pool) AddWorkers(amount int) {
	for id := 0; id < amount; id++ {
		w := NewWorker(id, p.Config.Client, p.Config.Process, p.WG)
		p.Workers = append(p.Workers, w)
		p.Outputs = append(p.Outputs, w.Output)
	}
	p.Sanitize()
	p.Deal()
}

// Connects outputs to Pool
//
// Reshuffles workers connections
func (p *Pool) Connect(outputs []chan []byte) {
	p.Inputs = append(p.Inputs, outputs...)
	p.Sanitize()
	p.Deal()
}

// Deals inputs to workers
func (p *Pool) Deal() {
	size := len(p.Workers)
	for i, input := range p.Inputs {
		id := i % size
		p.Workers[id].Inputs = append(p.Workers[id].Inputs, input)
	}
}

// Clears workers connections without deletion
func (p *Pool) Sanitize() {
	for _, w := range p.Workers {
		w.Inputs = []chan []byte{}
	}
}

// Disconnects all inputs
func (p *Pool) Disconnect() {
	p.Inputs = []chan []byte{}
	p.Sanitize()
}

// Checks if Pool is ready to Start
func (p *Pool) Ready() bool {
	for _, w := range p.Workers {
		if !w.Ready() {
			return false
		}
	}
	return true
}

// Starts Pool and PoolManager
func (p *Pool) Start(ctx context.Context) {
	p.Manager.Time.Begin = time.Now()
	go p.Manager.Start(ctx)
	for _, w := range p.Workers {
		go w.Start(ctx)
	}
	p.WG.Wait()
}
