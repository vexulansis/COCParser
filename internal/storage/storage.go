package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		DB: db,
	}
}

func GetDBURI() string {
	var DBURI string
	err := godotenv.Load()
	if err != nil {
		log.Fatal()
	}
	DBURI = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("DBPORT"),
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("DBNAME"),
	)
	return DBURI
}
func ConnectDB() *sql.DB {
	DBURI := GetDBURI()
	DB, err := sql.Open("postgres", DBURI)
	if err != nil {
		log.Fatal()
	}
	return DB
}
