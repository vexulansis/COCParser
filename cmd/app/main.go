package main

import (
	"fmt"
	"log"
	"time"

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
	n := 1000
	fmt.Println(DC, AC, CC)
	start := time.Now()
	pool := db.NewPool(5, DC)
	pool.Start()
	err = db.GenerateCredentials(n, DC)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	pool.WG.Wait()
	total := time.Since(start)
	fmt.Printf("Time to execute %d queries: %s", n, total)
}
