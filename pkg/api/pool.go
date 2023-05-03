package api

import (
	"sync"
)

type Pool struct {
	Size    int
	Workers []*Worker
	AC      *APIClient
	Mutex   *sync.Mutex
	WG      *sync.WaitGroup
}

func NewPool(size int, AC *APIClient) *Pool {
	pool := &Pool{
		Size:  size,
		AC:    AC,
		Mutex: &sync.Mutex{},
		WG:    &sync.WaitGroup{},
	}
	for i := 0; i < size; i++ {
		w := NewWorker(i, pool)
		pool.Workers = append(pool.Workers, w)
	}
	return pool
}
func NewPoolWithKeys(AC *APIClient) *Pool {
	pool := &Pool{
		Size:  len(AC.KeyPool),
		AC:    AC,
		Mutex: &sync.Mutex{},
		WG:    &sync.WaitGroup{},
	}
	for i, k := range AC.KeyPool {
		w := NewWorker(i, pool)
		w.Token = k.Key
		pool.Workers = append(pool.Workers, w)
	}
	return pool
}
func (p *Pool) Start() {
	for _, w := range p.Workers {
		go w.Start()
	}
}
func (p *Pool) Stop() {
	for _, w := range p.Workers {
		go w.Stop()
	}
}
