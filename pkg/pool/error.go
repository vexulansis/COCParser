package pool

import "fmt"

type ErrorHandler struct {
	// Cluster name
	Name string
	// Number of goroutines handling errors
	Size int
	// Manager for benchmarking
	Manager *ErrorManager
	// Inputs to listen
	Inputs []chan []byte
	// Channel to stop ErrorHandler
	Quit chan bool
}

func NewErrorHandler(name string, size int) *ErrorHandler {
	return &ErrorHandler{
		Name:    name,
		Size:    size,
		Manager: NewErrorManager(name),
		Inputs:  []chan []byte{},
		Quit:    make(chan bool),
	}
}
func (h *ErrorHandler) Prepare() {
	for id := 0; id < h.Size; id++ {
		input := make(chan []byte)
		h.Inputs = append(h.Inputs, input)
	}
}
func (h *ErrorHandler) Start() {
	for _, input := range h.Inputs {
		go h.HandleChan(input)
	}
}
func (h *ErrorHandler) HandleChan(input chan []byte) {
	for {
		select {
		case msg := <-input:
			fmt.Printf("msg: %v\n", msg)
		case <-h.Quit:
			return
		}
	}
}
