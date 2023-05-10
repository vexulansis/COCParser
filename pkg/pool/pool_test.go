package pool

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	. "github.com/vexulansis/COCParser/pkg/client"
	. "github.com/vexulansis/COCParser/pkg/task"
)

const testDelay = 1
const testInputsNum = 1000
const testWorkersNum = 7

var testFunc = func(msg []byte, client Client) ([]byte, error) {
	return nil, nil
}

func TestPool(t *testing.T) {
	testTask := &Task{
		Type: "TEST",
		Data: "DATA",
	}
	testMessage, _ := json.Marshal(testTask)
	testClient := &TestClient{
		HTTP: resty.New(),
	}
	testConfig := &PoolConfig{
		Client:  testClient,
		Process: testFunc,
		Delay:   testDelay,
	}
	testPool := NewPool("TEST", testConfig)
	testPool.AddWorkers(testWorkersNum)
	testInputs := []chan []byte{}
	for i := 0; i < testInputsNum; i++ {
		testInput := make(chan []byte, 100)
		testInputs = append(testInputs, testInput)
	}
	testPool.Connect(testInputs)
	testContext := context.Background()
	go testPool.Start(testContext)
	for _, input := range testInputs {
		time.Sleep(time.Millisecond * 10)
		input <- testMessage
	}
	time.Sleep(5 * time.Second)

	testContext.Done()
}
