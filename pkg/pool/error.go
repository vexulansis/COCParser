package pool

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type ErrorHandler struct {
	// Balancing number of goroutines by limitting channel buffer
	Size chan struct{}
	// Manager for benchmarking
	Manager *ErrorManager
	// Task counter
	WG *sync.WaitGroup
	// Input for errors
	Input chan []byte
}

// Wrapped error
type Error struct {
	Time time.Time
	Info string
}

// Creates ErrorHandler example
func NewErrorHandler(size int) *ErrorHandler {
	return &ErrorHandler{
		Size:    make(chan struct{}, size),
		Manager: NewErrorManager(),
		WG:      &sync.WaitGroup{},
		Input:   make(chan []byte),
	}
}

// Listening to input and printing Errors
func (h *ErrorHandler) Start() {
	for {
		select {
		case msg := <-h.Input:
			e := Error{}
			json.Unmarshal(msg, &e)
			fmt.Printf("Time: %s\nInfo: %s \n", e.Time.Format(timeformat), e.Info)
		default:
		}
	}
}

// Wraps error before sending to channel
func Wrap(err error) []byte {
	e := &Error{
		Time: time.Now(),
		Info: err.Error(),
	}
	msg, _ := json.Marshal(e)
	return msg
}
