package main

import (
	"log"

	tg "github.com/vexulansis/COCParser/internal/telegram"
)

// Точка входа
func main() {
	bot, err := tg.NewBot()
	if err != nil {
		log.Fatal()
	}
	err = bot.Start()
	if err != nil {
		log.Fatal()
	}
}
