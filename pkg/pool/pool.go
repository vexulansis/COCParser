package pool

import (
	"database/sql"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-resty/resty/v2"
)

type Pool struct {
	// Just to identify
	Name string
	// Balancing Pool size by limitting channel buffer
	Size int
	// Manager for benchmarking
	Manager *PoolManager
	// Database pointer to send queries
	DB *sql.DB
	// http.Client pointer to send requests
	Client *resty.Client
	// Current state
	Current atomic.Uint64
	// Tracking workers
	WG *sync.WaitGroup
	// Inputs to connect
	Inputs []chan []byte
	// ErrorHanlder channel
	Error chan []byte
	// WorkerPool
	Workers []*Worker
}

func NewPool(name string, size int) *Pool {
	return &Pool{
		Name:    name,
		Size:    size,
		Manager: NewPoolManager(name),
		WG:      &sync.WaitGroup{},
		Workers: []*Worker{},
	}
}
func (p *Pool) ConnectGenerator(generator *Generator) {
	p.Inputs = generator.Outputs
}
func (p *Pool) Prepare() {
	for id := 0; id < p.Size; id++ {
		w := NewWorker(id, p)
		p.Workers = append(p.Workers, w)
	}
}
func (p *Pool) Start() {
	p.Manager.Time.Begin = time.Now()
	p.WG.Add(len(p.Workers))
	for _, w := range p.Workers {
		go w.Start()
	}
}
func (p *Pool) Stop() {
	for _, w := range p.Workers {
		go w.Stop()
	}
	p.Manager.Time.End = time.Now()
}
