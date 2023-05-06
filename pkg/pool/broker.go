package pool

import (
	"sync"
)

type Broker struct {
	// Mutex to avoid concurrent map RW
	Mutex *sync.Mutex
	// List of input channels
	Inputs []*Conn
	// List of output channels
	Outputs []*Conn
	// List of error channels
	Errors []*Conn
	// Staight path for fast access
	ErrorHandler chan []byte
	// Map type [inputID]: [outputID...]
	Config map[int][]int
	// Channel to stop the Broker
	Quit chan bool
}
type Conn struct {
	Port    int
	Channel chan []byte
}

// Creating new Broker with connection to ErrorHandler
func NewBroker(errorHandler chan []byte) *Broker {
	return &Broker{
		Mutex:        &sync.Mutex{},
		Inputs:       []*Conn{},
		Outputs:      []*Conn{},
		Errors:       []*Conn{},
		ErrorHandler: errorHandler,
		Config:       make(map[int][]int),
		Quit:         make(chan bool),
	}
}

// Starts Broker
func (b *Broker) Start() {
	for {
		errors := b.Errors
		// Check if any error channels available
		if len(errors) > 0 {
			for _, err := range errors {
				select {
				case msg := <-err.Channel:
					b.ErrorHandler <- msg
				case <-b.Quit:
					return
				}
			}
		}
	}
}

// Stops operating and benchmarking
func (b *Broker) Stop() {
	b.Quit <- true
}

// Connecting to Broker input, returning port
func (b *Broker) ConnectInput(input chan []byte) int {
	port := len(b.Inputs)
	conn := &Conn{
		Port:    port,
		Channel: input,
	}
	b.Inputs = append(b.Inputs, conn)
	// Mutex to avoid concurrent map RW
	b.Mutex.Lock()
	b.Config[port] = []int{}
	b.Mutex.Unlock()
	go b.RouteInput(port)
	return port
}

// Connecting to Broker output, returning port
func (b *Broker) ConnectOutput(output chan []byte) int {
	port := len(b.Outputs)
	conn := &Conn{
		Port:    port,
		Channel: output,
	}
	b.Outputs = append(b.Outputs, conn)
	return port
}

// Connecting error channel, fan in into ErrorHandler
func (b *Broker) ConnectError(err chan []byte) int {
	port := len(b.Errors)
	conn := &Conn{
		Port:    port,
		Channel: err,
	}
	b.Errors = append(b.Errors, conn)
	return port
}

// Connecting input and output
func (b *Broker) Bridge(inputId int, outputId int) {
	// Mutex to avoid concurrent map RW
	b.Mutex.Lock()
	b.Config[inputId] = append(b.Config[inputId], outputId)
	b.Mutex.Unlock()
}

// Listening to input and resending messages to outputs
func (b *Broker) RouteInput(inputID int) {
	for {
		// Mutex to avoid concurrent map RW
		b.Mutex.Lock()
		outputs := b.Config[inputID]
		b.Mutex.Unlock()
		// Check if any outputs available
		if len(outputs) > 0 {
			select {
			case msg := <-b.Inputs[inputID].Channel:
				for _, outputID := range outputs {
					b.Outputs[outputID].Channel <- msg
				}
			case <-b.Quit:
				return
			}
		}
	}
}
