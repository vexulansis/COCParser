package db

import (
	"testing"
)

const expected = 100000

var received int
var extract bool

func TestGenerator(t *testing.T) {
	taskChan := make(chan *Task)
	reserved := &Reserved{
		Total:    expected,
		Email:    "example+%d@email.com",
		Password: "qwerty",
	}
	g := NewGenerator(taskChan)
	go g.GenerateReserved(reserved)
	for {
		select {
		case _, ok := <-taskChan:
			received++
			if !ok {
				extract = true
			}
		}
		if extract == true {
			break
		}
	}
	if received != expected+1 {
		t.Errorf("Expected: %d , Received: %d", expected, received)
	}
	g.Manager.PrintBenchmark()
}
