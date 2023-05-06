package pool

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type ErrorHandler struct {
	Name string
	// Balancing number of goroutines by limitting channel buffer
	Size chan struct{}
	// Manager for benchmarking
	Manager *ErrorManager
	// Custom printing features
	Logger *Logger
	// Task counter
	WG *sync.WaitGroup
	// Input for errors
	Input chan []byte
	// Channel to stop ErrorHandler
	Quit chan bool
}

// Creates ErrorHandler example
func NewErrorHandler(name string, size int) *ErrorHandler {
	return &ErrorHandler{
		Size:    make(chan struct{}, size),
		Manager: NewErrorManager(name),
		Logger:  NewLogger(),
		WG:      &sync.WaitGroup{},
		Input:   make(chan []byte),
		Quit:    make(chan bool),
	}
}

// Listening to input and printing Errors
func (h *ErrorHandler) Start() {
	h.Manager.Time.Begin = time.Now()
	for {
		select {
		case msg := <-h.Input:
			h.Manager.Received.Add(1)
			h.Logger.Error(QuickSprint(msg))
		case <-h.Quit:
			return
		}
	}
}
func (h *ErrorHandler) Stop() {
	h.Manager.Time.End = time.Now()
	h.Quit <- true
}

// Custom error type
type Error struct {
	Time time.Time
	Info string
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

// Unwraps error before printing
func Unwrap(msg []byte) *Error {
	e := &Error{}
	json.Unmarshal(msg, e)
	return e
}

// Print Error
func (e *Error) Print() {
	fmt.Printf("Time: %s\nInfo: %s \n", e.Time.Format(timeformat), e.Info)
}

// Converts Error to string
func (e *Error) Sprint() string {
	return fmt.Sprintf("Time: %s\nInfo: %s \n", e.Time.Format(timeformat), e.Info)
}

// Combines Unwrap and Print
func QuickPrint(msg []byte) {
	Unwrap(msg).Print()
}

// Combines Unwrap and Sprint
func QuickSprint(msg []byte) string {
	return Unwrap(msg).Sprint()
}
