package main

import (
	"github.com/vexulansis/COCParser/internal/db"
	"github.com/vexulansis/COCParser/pkg/api"
)

func main() {
	logger := InitMainLogger()
	DC, err := db.NewClient()
	if err != nil {
		f := MainLoggerFields{
			Source:  "DBCLIENT",
			Method:  "CREATION",
			Subject: "ERROR",
		}
		logger.Fatal(f, err)
	}
	AC, err := api.NewClient()
	if err != nil {
		f := MainLoggerFields{
			Source:  "APICLIENT",
			Method:  "CREATION",
			Subject: "ERROR",
		}
		logger.Fatal(f, err)
	}
	Prepare(AC, DC, logger)
}
