package db

import "fmt"

type Worker struct {
	ID   int
	Pool *Pool
	Quit chan bool
}

func NewWorker(ID int, pool *Pool) *Worker {
	return &Worker{
		ID:   ID,
		Pool: pool,
		Quit: make(chan bool),
	}
}
func (w *Worker) Start() {
	f := DBLoggerFields{
		Source:      fmt.Sprintf("WORKER#%d", w.ID),
		Method:      "START",
		Subject:     "<---",
		Destination: "TASKCHANNEL",
	}
	w.Pool.DC.Logger.Print(f, 0)
	for {
		select {
		case task := <-w.Pool.DC.TaskChan:
			w.Pool.WG.Add(1)
			err := w.Process(task)
			if err != nil {

			}
		case <-w.Quit:
			return
		}
	}
}
func (w *Worker) Process(task *Task) error {
	defer w.Pool.WG.Done()
	switch t := task.Data.(type) {
	case Credentials:
		res, err := w.Pool.DC.DB.Exec("insert into credentials(email,password) values ($1,$2)", t.Email, t.Password)
		if err != nil {
			return err
		}
		f := DBLoggerFields{
			Source:      fmt.Sprintf("WORKER#%d", w.ID),
			Method:      "INSERT",
			Subject:     t.Email,
			Destination: "credentials",
		}
		w.Pool.DC.Logger.Print(f, res)
	}
	return nil
}
func (w *Worker) Stop() {
	f := DBLoggerFields{
		Source:      fmt.Sprintf("WORKER#%d", w.ID),
		Method:      "STOP",
		Subject:     "<---",
		Destination: "TASKCHANNEL",
	}
	w.Pool.DC.Logger.Print(f, 0)
	go func() {
		w.Quit <- true
	}()
}
