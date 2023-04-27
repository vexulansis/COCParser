package main

import (
	"github.com/vexulansis/COCParser/internal/storage"
	api "github.com/vexulansis/COCParser/pkg/api"
)

func main() {
	storage.NewStorage()
	api.NewClient()
}
