package db

import (
	"sync"
	"time"
)

type Pool struct {
	Size    int
	Workers []*Worker
	DC      *DBClient
	Manager *PoolManager
	WG      *sync.WaitGroup
	Input   <-chan *Task
	Output  chan<- any
}

func NewPool(size int, input <-chan *Task, output chan<- any, DC *DBClient) *Pool {
	m := initPoolManager()
	pool := &Pool{
		Size:    size,
		DC:      DC,
		Manager: m,
		WG:      &sync.WaitGroup{},
		Input:   input,
		Output:  output,
	}
	for i := 0; i < size; i++ {
		w := NewWorker(i, pool)
		pool.Workers = append(pool.Workers, w)
	}
	return pool
}
func (p *Pool) Start() {
	p.Manager.StartTime = time.Now()
	for _, w := range p.Workers {
		go w.Start()
	}
}
func (p *Pool) Stop() {
	p.Manager.ShutdownTime = time.Now()
	for _, w := range p.Workers {
		w.Stop()
	}
}
