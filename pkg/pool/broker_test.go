package pool

import (
	"errors"
	"fmt"
	"testing"
)

func TestBroker(t *testing.T) {
	handler := NewErrorHandler(10)
	go handler.Start()
	broker := NewBroker(handler.Input)
	go broker.Start()
	testChan := make(chan []byte)
	broker.ConnectError(testChan)
	for id := 0; id < 100; id++ {
		testChan <- Wrap(errors.New(fmt.Sprintf("TEST %d", id)))
	}
}
