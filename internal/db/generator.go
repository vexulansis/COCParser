package db

import (
	"fmt"
	"sync"
	"time"
)

type Generator struct {
	Manager *GeneratorManager
	WG      *sync.WaitGroup
	Output  chan<- *Task
}
type Reserved struct {
	Total    int
	Email    string
	Password string
}

func NewGenerator(taskChan chan<- *Task) *Generator {
	m := initGeneratorManager()
	return &Generator{
		Manager: m,
		WG:      &sync.WaitGroup{},
		Output:  taskChan,
	}
}
func (g *Generator) GenerateReserved(reserved *Reserved) {
	for id := 0; id < reserved.Total; id++ {
		g.WG.Add(1)
		go func(id int, output chan<- *Task) {
			defer g.WG.Done()
			// Creating account example
			acc := &Account{
				ID: id,
				Credentials: Credentials{
					Email:    fmt.Sprintf(reserved.Email, id),
					Password: reserved.Password,
				},
			}
			// Creating task
			task := &Task{
				ID:   id,
				Type: "INSERT",
				Data: acc,
			}
			// Sending task
			output <- task
			// Counting tasks
			g.Manager.Mutex.Lock()
			g.Manager.TasksGenerated++
			g.Manager.Mutex.Unlock()
		}(id, g.Output)
	}
	g.Manager.ShutdownTime = time.Now()
	g.WG.Wait()
	close(g.Output)
}
