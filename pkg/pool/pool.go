package pool

import (
	"database/sql"
	"sync"
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
	// WorkerPool
	Workers []*Worker
	// Task counter
	WG *sync.WaitGroup
	// Channel with Broker port ID to listen to
	Input *Conn
	// Channel with Broker port ID to send results
	Output *Conn
	// Channel with Broker port ID to send errors
	Errors *Conn
	// Identify if Pool is ready to operate
	Ready bool
	// Channel to stop the Pool
	Quit chan bool
}

// New Pool example, call Prepare() before Start()
func NewPool(name string, size int) *Pool {
	pool := &Pool{
		Name: name,
		Size: size,
		WG:   &sync.WaitGroup{},
		Input: &Conn{
			Channel: make(chan []byte),
		},
		Output: &Conn{
			Channel: make(chan []byte),
		},
		Errors: &Conn{
			Channel: make(chan []byte),
		},
		Quit: make(chan bool),
	}
	return pool
}

// Initializing Manager, connecting to Broker, creating Workers
func (p *Pool) Prepare(broker *Broker) {
	// Initializing Manager
	p.Manager = NewPoolManager()
	// Connecting to Broker
	p.Input.Port = broker.ConnectOutput(p.Input.Channel)
	p.Output.Port = broker.ConnectInput(p.Output.Channel)
	// Creating workers
	for id := 0; id < p.Size; id++ {
		w := NewWorker(id, p)
		p.Workers = append(p.Workers, w)
	}
}

// Starts operating and benchmarking
func (p *Pool) Start() {
	p.Manager.Time.Begin = time.Now()
	for _, w := range p.Workers {
		go w.Start()
	}
}

// Stops operating and benchmarking
func (p *Pool) Stop() {
	p.Manager.Time.End = time.Now()
	for _, w := range p.Workers {
		go w.Stop()
	}
}

// Wrapping and sending errors
func (p *Pool) Error(err error) {
	p.Errors.Channel <- Wrap(err)
}
