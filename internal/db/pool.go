package db

import (
	"fmt"
	"sync"
)

type Pool interface {
	Start()
	Stop()
	AddWork(Task)
}
type WorkerPool struct {
	DC        *DBClient
	Size      int
	Tasks     chan Task
	WG        *sync.WaitGroup
	StartOnce sync.Once
	StopOnce  sync.Once
	Quit      chan struct{}
}

func NewWorkerPool(size int, channelsize int, DC *DBClient) *WorkerPool {
	tasks := make(chan Task, channelsize)
	quit := make(chan struct{})
	return &WorkerPool{
		DC:        DC,
		Size:      size,
		Tasks:     tasks,
		WG:        &sync.WaitGroup{},
		StartOnce: sync.Once{},
		StopOnce:  sync.Once{},
		Quit:      quit,
	}
}
func (wp *WorkerPool) Start() {
	wp.StartOnce.Do(func() {
		f := DBLoggerFields{
			Source: "WORKERPOOL",
			Method: "START",
		}
		wp.DC.Logger.Print(f, 0)
		wp.startWorkers()
	})
}
func (wp *WorkerPool) Stop() {
	wp.StopOnce.Do(func() {
		f := DBLoggerFields{
			Source: "WORKERPOOL",
			Method: "STOP",
		}
		wp.DC.Logger.Print(f, 0)
		close(wp.Quit)
	})
}
func (wp *WorkerPool) startWorkers() {
	for id := 0; id < wp.Size; id++ {
		wp.WG.Add(1)
		go func(workerNum int) {
			f := DBLoggerFields{
				Source: fmt.Sprintf("WORKER#%d", id),
				Method: "START",
			}
			wp.DC.Logger.Print(f, 0)
			for {
				select {
				case <-wp.Quit:
					f := DBLoggerFields{
						Source: fmt.Sprintf("WORKER#%d", id),
						Method: "STOP",
					}
					wp.DC.Logger.Print(f, 0)
					return
				case task, ok := <-wp.Tasks:
					if !ok {
						f := DBLoggerFields{
							Source: fmt.Sprintf("WORKER#%d", id),
							Method: "STOP",
						}
						wp.DC.Logger.Print(f, 0)
						return
					}

					task.Execute(wp.DC, id)
				}
			}
		}(id)
	}
}
func (wp *WorkerPool) AddTask(t Task) {
	select {
	case wp.Tasks <- t:
	case <-wp.Quit:
	}
}
