package pool

import (
	"encoding/json"
	"testing"
)

func TestPool(t *testing.T) {
	handler := NewErrorHandler("test", 10)
	go handler.Start()
	broker := NewBroker(handler.Input)
	go broker.Start()
	pool := NewPool("test", 25)
	pool.Prepare(broker)
	testChan := make(chan []byte)
	testID := broker.ConnectInput(testChan)
	broker.Bridge(testID, pool.Input.Port)
	go pool.Start()
	for id := 0; id < 100; id++ {
		testMsg := &Message{
			ID:     id,
			Client: "TEST",
			Type:   "TEST",
			Data:   "TEST",
		}
		msg, _ := json.Marshal(testMsg)
		testChan <- msg
	}
	pool.Stop()
	broker.Stop()
	handler.Stop()
	handler.Manager.Stats()
	pool.Manager.Stats()
}
