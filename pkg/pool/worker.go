package pool

import (
	"encoding/json"
	"errors"
)

// Unit representing goroutine
type Worker struct {
	ID   int
	Pool *Pool
	Quit chan bool
}

// New Worker example
func NewWorker(id int, pool *Pool) *Worker {
	return &Worker{
		ID:   id,
		Pool: pool,
		Quit: make(chan bool),
	}
}

// Starts listening to channel and processing tasks
func (w *Worker) Start() {
	for {
		select {
		case msg := <-w.Pool.Input.Channel:
			w.Pool.Manager.Received.Add(1)
			w.Pool.WG.Add(1)
			err := w.Process(msg)
			if err != nil {
				w.Pool.Error(err)
			}
		case <-w.Quit:
			return
		}
	}
}

// Task processing logic core
func (w *Worker) Process(msg []byte) error {
	defer w.Pool.WG.Done()
	defer w.Pool.Manager.Processed.Add(1)
	task := Message{}
	err := json.Unmarshal(msg, &task)
	if err != nil {
		return err
	}
	// Switch between possible task clients
	switch task.Client {
	case "DB":
		w.Query(task)
	case "HTTP":
		w.Request(task)
	default:
		return errors.New("Incorrect task client")
	}
	return nil
}

// Stops operating
func (w *Worker) Stop() {
	w.Quit <- true
}
