package pool

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type Generator struct {
	// Just to identify
	Name string
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
	// Channel to stop the Generator
	Quit chan bool
}

func NewGenerator(name string, size int) *Generator {
	return &Generator{
		Name: name,
		Size: make(chan struct{}, size),
		WG:   &sync.WaitGroup{},
		Output: &Conn{
			Channel: make(chan []byte),
		},
		Errors: &Conn{
			Channel: make(chan []byte),
		},
		Quit: make(chan bool),
	}
}

// Initializing Manager and connecting to Broker
func (g *Generator) Prepare(broker *Broker) {
	// Initializing Manager
	g.Manager = NewGeneratorManager(g.Name)
	// Connecting to Broker
	g.Output.Port = broker.ConnectInput(g.Output.Channel)
	g.Errors.Port = broker.ConnectError(g.Errors.Channel)
}

// Generator logic core
func (g *Generator) Generate(example any, amount int) {
	g.Manager.Time.Begin = time.Now()
	g.WG.Add(amount)
	switch data := example.(type) {
	case Account:
		for id := 0; id < amount; id++ {
			// Blocks goroutine creation if buffer is full
			g.Size <- struct{}{}
			go func(id int, data Account) {
				defer g.WG.Done()
				data.ID = id
				data.Email = fmt.Sprintf(data.Email, id)
				task := &Message{
					ID:     id,
					Client: "DB",
					Type:   "INSERTACCOUNT",
					Data:   data,
				}
				msg, err := json.Marshal(task)
				if err != nil {
					g.Errors.Channel <- Wrap(err)
				}
				g.Output.Channel <- msg
				g.Manager.Generated.Add(1)
				// Unblocks goroutine creation when done
				<-g.Size
			}(id, data)
		}
	case int:

	}
	g.Manager.Time.End = time.Now()
}
