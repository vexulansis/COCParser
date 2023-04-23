package main

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/vexulansis/COCParser/internal/storage"
	"github.com/vexulansis/COCParser/pkg/api"
)

func main() {
	db := storage.NewStorage()
	client := api.Client{
		Storage: db}
	_, err := toml.DecodeFile("keys.toml", &client)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	client.UpdateClans(0, 2950)
	time.Sleep(time.Hour * 1)
}
