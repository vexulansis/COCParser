package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/vexulansis/COCParser/internal/db"
	. "github.com/vexulansis/COCParser/pkg/pool"
)

func main() {
	example := Account{}
	DC := db.NewClient()
	_, err := toml.DecodeFile("reference.toml", &example)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		log.Fatal()
	}
	//
	handler := NewErrorHandler("main", 1)
	go handler.Start()
	//
	broker := NewBroker(handler.Input)
	go broker.Start()
	//
	pool := NewPool("main", 10)
	pool.ConnectDB(DC.DB)
	pool.Prepare(broker)
	//
	generator := NewGenerator("accounts", 1000)
	generator.Prepare(broker)
	//
	broker.Bridge(generator.Output.Port, pool.Input.Port)
	go pool.Start()
	generator.Generate(example, 100000)
	//
	generator.WG.Wait()
	generator.Manager.Stats()
	pool.Stop()
	broker.Stop()
	handler.Stop()
	handler.Manager.Stats()
	pool.Manager.Stats()
}
