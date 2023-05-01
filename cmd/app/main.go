package main

import (
	"fmt"

	"github.com/vexulansis/COCParser/internal/storage"
	"github.com/vexulansis/COCParser/pkg/api"
)

func main() {
	ClientDB, err := storage.NewClient()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	// err = ClientDB.GenerateAccounts(100)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	_, err = api.NewClient(ClientDB)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	// err = ClientAPI.SanitizeKeys()
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }
	// err = ClientAPI.FillKeys()
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }
}
