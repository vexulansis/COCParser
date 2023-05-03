package db

import (
	"fmt"
	"sync"
)

type Pool struct {
	Size    int
	Workers []*Worker
	DC      *DBClient
	WG      *sync.WaitGroup
}

func NewPool(size int, DC *DBClient) *Pool {
	pool := &Pool{
		Size: size,
		DC:   DC,
		WG:   &sync.WaitGroup{},
	}
	for i := 0; i < size; i++ {
		w := NewWorker(i, pool)
		pool.Workers = append(pool.Workers, w)
		fmt.Printf("w: %v\n", w)
	}
	return pool
}
func (p *Pool) Start() {
	for _, w := range p.Workers {
		go w.Start()
	}
}
