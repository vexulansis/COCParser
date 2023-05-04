package db

import "testing"

const size = 10
const tasks = 10000000

func TestPool(t *testing.T) {
	input := make(chan *Task)
	output := make(chan any)
	DC, _ := NewClient()
	pool := NewPool(100, input, output, DC)
	pool.Start()
	for i := 0; i < tasks; i++ {
		pool.WG.Add(1)
		task := &Task{
			ID:   i,
			Type: "TEST",
		}
		input <- task
	}
	pool.WG.Wait()
	pool.Stop()
	pool.Manager.PrintBenchmark()
}
