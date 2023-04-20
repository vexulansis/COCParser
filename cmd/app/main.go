package main

import (
	"fmt"
	"log"

	alg "github.com/vexulansis/COCParser/internal/algo"
	db "github.com/vexulansis/COCParser/internal/storage"
)

var refstr = "0289CGJLPQRUVYO"

// Точка входа
func main() {
	// bot, err := tg.NewBot()
	// if err != nil {
	// 	log.Fatal()
	// }
	storage := db.NewStorage()
	enum := alg.NewEnumerator(refstr, storage)
	err := enum.Execute()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		log.Fatal()
	}
	// err = bot.Start()
	// if err != nil {
	// 	log.Fatal()
	// }
}
