package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	. "github.com/vexulansis/COCParser/pkg/client"
	. "github.com/vexulansis/COCParser/pkg/task"
)

const testID = 12

var testFunc = func(msg []byte, client Client) ([]byte, error) {
	testTask := &Task{}
	json.Unmarshal(msg, testTask)
	fmt.Printf("%v %v IN PROCESS\n", testTask.Type, testTask.Data)
	client.Exec(*testTask)
	return nil, nil
}

func TestWorker(t *testing.T) {
	testTask := &Task{
		Type: "TEST",
		Data: "DATA",
	}
	testClient := &TestClient{
		HTTP: resty.New(),
	}
	testWG := &sync.WaitGroup{}
	testWorker := NewWorker(testID, testClient, testFunc, testWG)
	testContext := context.Background()
	testInput := make(chan []byte)
	testMessage, _ := json.Marshal(testTask)
	//
	testWorker.AddInput(testInput)
	go testWorker.Start(testContext)
	testInput <- testMessage
	time.Sleep(time.Second * 10)
	testContext.Done()
}
