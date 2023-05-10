package generator

import (
	"context"
	"encoding/json"
	"sync"
	"sync/atomic"
	"time"

	. "github.com/vexulansis/COCParser/pkg/task"
)

var ref = "0289PYLQGRJCUV"

type GenerateFunc func(output chan []byte, counter *atomic.Uint64, goal uint64, wg *sync.WaitGroup, example any)

type Generator struct {
	Name    string
	Config  *GeneratorConfig
	Manager *GeneratorManager
	WG      *sync.WaitGroup
	Outputs []chan []byte
}
type GeneratorConfig struct {
	Size     int
	Generate GenerateFunc
	Example  any
	Delay    int
}

func NewGenerator(name string, config *GeneratorConfig) *Generator {
	generator := &Generator{
		Name:   name,
		Config: config,
		WG:     &sync.WaitGroup{},
	}
	for i := 0; i < config.Size; i++ {
		output := make(chan []byte, 100000)
		generator.Outputs = append(generator.Outputs, output)

	}
	generator.Manager = NewGeneratorManager(generator)
	return generator
}
func (g *Generator) Start(ctx context.Context, amount uint64) {
	g.Manager.Time.Begin = time.Now()
	go g.Manager.Start(ctx)
	for _, output := range g.Outputs {
		g.WG.Add(1)
		go g.Config.Generate(output, &g.Manager.Generated, amount, g.WG, g.Config.Example)
	}
	g.WG.Wait()
}
func GenerateTags(output chan []byte, counter *atomic.Uint64, goal uint64, wg *sync.WaitGroup, example any) {
	defer wg.Done()
	testTask := &Task{
		Type: "GETCLAN",
	}
	for {
		current := counter.Add(1)
		if current > goal {
			return
		}
		tag := getTagFromID(int(current))
		testTask.Data = tag
		msg, _ := json.Marshal(testTask)
		output <- msg
	}

}
func getTagFromID(id int) string {
	var index int
	var tag string
	size := len(ref)
	for id > 0 {
		index = id % size
		tag = string(ref[index]) + tag
		id -= index
		id /= size
	}
	return tag
}
