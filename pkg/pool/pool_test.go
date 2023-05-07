package pool

import (
	"testing"
)

func TestPool(t *testing.T) {
	//
	pool := NewPool("TEST", 1000)
	pool.Prepare()
	//
	generator := NewGenerator("TEST", 1000)
	generator.Prepare()
	//
	pool.ConnectGenerator(generator)
	generator.Start(0, 10000000)
	pool.Start()
	generator.WG.Wait()
	generator.Stop()
	pool.Stop()
	pool.Manager.Stats()
	//
}
