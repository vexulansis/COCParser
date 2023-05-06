package pool

import (
	"testing"
)

func TestGenerator(t *testing.T) {
	example := Account{
		Email:    "test%s@example.com",
		Password: "qwerty",
	}
	//
	handler := NewErrorHandler("test", 10)
	go handler.Start()
	//
	broker := NewBroker(handler.Input)
	go broker.Start()
	//
	pool := NewPool("test", 25)
	pool.Prepare(broker)
	//
	generator := NewGenerator("test", 30)
	generator.Prepare(broker)
	//
	broker.Bridge(generator.Output.Port, pool.Input.Port)
	go pool.Start()
	generator.Generate(example, 1000)
	//
	pool.WG.Wait()
	generator.WG.Wait()
	generator.Manager.Stats()
	pool.Stop()
	broker.Stop()
	handler.Stop()
	handler.Manager.Stats()
	pool.Manager.Stats()
}
