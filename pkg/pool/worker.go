package pool

import (
	"encoding/json"
)

// Unit representing goroutine
type Worker struct {
	ID     int
	Pool   *Pool
	Output chan []byte
	Quit   chan bool
}

func NewWorker(id int, pool *Pool) *Worker {
	return &Worker{
		ID:     id,
		Pool:   pool,
		Output: make(chan []byte),
		Quit:   make(chan bool),
	}
}
func (w *Worker) Start() {
	defer w.Pool.WG.Done()
	for {
		current := w.Pool.Current.Add(1)
		inputID := int(current) % len(w.Pool.Inputs)
		input := w.Pool.Inputs[inputID]
		select {
		case msg := <-input:
			w.Pool.Manager.Received.Add(1)
			w.Process(msg)
		case <-w.Quit:
			return
		default:
			continue
		}
	}
}

func (w *Worker) Process(msg []byte) error {
	defer w.Pool.Manager.Processed.Add(1)
	message := &Message{}
	json.Unmarshal(msg, message)
	return nil
}
func (w *Worker) Stop() {
	w.Quit <- true
}
