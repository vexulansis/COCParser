package pool

import (
	"encoding/json"
	"sync"
	"sync/atomic"
	"time"
)

type Generator struct {
	// Just to identify
	Name string
	// Balancing number of goroutines by limitting channel buffer
	Size int
	// Manager for benchmarking
	Manager *GeneratorManager
	// WaitGroup
	WG *sync.WaitGroup
	// Current state
	Current atomic.Uint64
	// Goal to reach
	Goal atomic.Uint64
	// To send errors
	Errors chan []byte
	// Outputs to send data
	Outputs []chan []byte
	// Channel to stop Generator
	Quit chan bool
}

func NewGenerator(name string, size int) *Generator {
	return &Generator{
		Name:    name,
		Size:    size,
		Manager: NewGeneratorManager(name),
		WG:      &sync.WaitGroup{},
		Outputs: []chan []byte{},
		Errors:  make(chan []byte),
		Quit:    make(chan bool, size),
	}
}
func (g *Generator) Prepare() {
	for id := 0; id < g.Size; id++ {
		output := make(chan []byte, 1000)
		g.Outputs = append(g.Outputs, output)
	}
}
func (g *Generator) Start(start uint64, end uint64) {
	g.Manager.Time.Begin = time.Now()
	g.Current.Store(start)
	g.Goal.Store(end)
	g.WG.Add(len(g.Outputs))
	for i, output := range g.Outputs {
		go g.GenerateToChan(output, i)
	}
}
func (g *Generator) Stop() {
	g.Manager.Time.End = time.Now()
	g.Manager.Stats()
}

func (g *Generator) GenerateToChan(output chan []byte, id int) {
	defer g.WG.Done()
	for {
		current := g.Current.Add(1)
		if current <= g.Goal.Load() {
			tag := getTagFromID(int(current))
			msg, _ := WrapTag(tag)
			output <- msg
			g.Manager.Generated.Add(1)
		} else {
			break
		}
	}
}

func WrapTag(tag string) ([]byte, error) {
	message := &Message{
		Client: "DB",
		Type:   "GETCLAN",
		Data:   tag,
	}
	msg, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
