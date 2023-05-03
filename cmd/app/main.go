package main

import (
	"fmt"
	"log"

	"github.com/vexulansis/COCParser/internal/cluster"
	"github.com/vexulansis/COCParser/internal/db"
	"github.com/vexulansis/COCParser/pkg/api"
)

func main() {

	DC, err := db.NewClient()
	if err != nil {
		log.Fatal()
	}
	AC, err := api.NewClient()
	if err != nil {
		log.Fatal()
	}
	CC, err := cluster.NewClient()
	if err != nil {
		log.Fatal()
	}
	fmt.Println(DC, AC, CC)
	wp := db.NewWorkerPool(10, 100, DC)
	wp.Start()
	err = db.GenerateCredentials(100, wp)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	wp.Stop()
}
