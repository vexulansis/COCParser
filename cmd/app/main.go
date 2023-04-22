package main

import db "github.com/vexulansis/COCParser/internal/storage"

var refstr = "0289CGJLPQRUVY"

// Точка входа
func main() {
	storage := db.NewStorage()
	enum := db.NewEnumerator(refstr, storage)
	enum.Storage.GenerateTags(2954)
	enum.Storage.AddClanInfo(2954)
}
