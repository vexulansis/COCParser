package pool

import (
	"sync"
)

type Broker struct {
	Mutex        *sync.Mutex
	Inputs       []*Conn
	Outputs      []*Conn
	Errors       []*Conn
	ErrorHandler chan []byte
	Config       map[int][]int
}
type Conn struct {
	Port    int
	Channel chan []byte
}

// Creating new Broker with connection to ErrorHandler
func NewBroker(errorHandler chan []byte) *Broker {
	return &Broker{
		ErrorHandler: errorHandler,
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
				default:

				}
			}
		}
	}
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
func (b *Broker) ConnectError(err chan []byte) {
	port := len(b.Errors)
	conn := &Conn{
		Port:    port,
		Channel: err,
	}
	b.Errors = append(b.Errors, conn)
	return
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
			default:
				return
			}
		}
	}
}
