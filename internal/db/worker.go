package db

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
	for {
		select {
		case task := <-w.Pool.Input:
			w.Pool.Manager.Mutex.Lock()
			w.Pool.Manager.TasksReceived++
			w.Pool.Manager.Mutex.Unlock()
			w.Process(task)
		case <-w.Quit:
			return
		}
	}
}
func (w *Worker) Process(task *Task) {
	defer w.Pool.WG.Done()
	switch task.Type {
	case "TEST":

	}
	w.Pool.Manager.Mutex.Lock()
	w.Pool.Manager.TasksProcessed++
	w.Pool.Manager.Mutex.Unlock()
}
func (w *Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}
