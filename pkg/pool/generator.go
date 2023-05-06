package pool

import (
	"sync"
)

type Generator struct {
	// Just to identify
	ID int
	// Balancing number of goroutines by limitting channel buffer
	Size chan struct{}
	// Manager for benchmarking
	Manager *GeneratorManager
	// Task counter
	WG *sync.WaitGroup
	// Channel with Broker port ID to send results
	Output *Conn
	// Channel with Broker port ID to send errors
	Errors *Conn
}
