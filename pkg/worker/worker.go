package worker

import (
	"reflect"
	"sync"

	. "github.com/vexulansis/COCParser/pkg/client"
	"golang.org/x/net/context"
)

type ProcessFunc func(msg []byte, client Client) ([]byte, error)

type Worker struct {
	ID      int
	Manager *WorkerManager
	Client  Client
	Process ProcessFunc
	WG      *sync.WaitGroup
	Inputs  []chan []byte
	Output  chan []byte
}

// Creates Worker example
func NewWorker(id int, client Client, process ProcessFunc, wg *sync.WaitGroup) *Worker {
	worker := &Worker{
		ID:      id,
		Client:  client,
		Process: process,
		WG:      wg,
		Output:  make(chan []byte, 1000),
	}
	worker.Manager = NewWorkerManager(worker)
	return worker
}

// Checks if Worker is ready to Start
func (w *Worker) Ready() bool {
	return len(w.Inputs) > 0
}

// Adds input channel
func (w *Worker) AddInput(input chan []byte) {
	w.Inputs = append(w.Inputs, input)
}

// Creates an array of cases to switch between
func (w *Worker) MakeCases() []reflect.SelectCase {
	cases := []reflect.SelectCase{}
	// Adding default case to continue operating
	c := reflect.SelectCase{Dir: reflect.SelectDefault}
	cases = append(cases, c)
	for _, input := range w.Inputs {
		c := reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(input)}
		cases = append(cases, c)
	}
	return cases
}

// Starts listenning to Inputs and processing messages
func (w *Worker) Start(ctx context.Context) {
	defer w.WG.Done()
	w.WG.Add(1)
	cases := w.MakeCases()
	for {
		switch index, msg, ok := reflect.Select(cases); ok {
		case index != 0:
			w.Manager.Received.Add(1)
			w.Process(msg.Bytes(), w.Client)
			w.Manager.Processed.Add(1)
		default:
			select {
			case <-ctx.Done():
				return
			default:
				continue
			}
		}

	}
}
