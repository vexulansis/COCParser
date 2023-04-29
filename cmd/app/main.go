package main

import (
	"fmt"

	"github.com/vexulansis/COCParser/internal/storage"
	api "github.com/vexulansis/COCParser/pkg/api"
)

func main() {
	db, err := storage.NewStorage()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	_, err = api.NewClient(db)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

}
