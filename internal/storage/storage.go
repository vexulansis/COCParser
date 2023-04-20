package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage() *Storage {
	db := ConnectDB()
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
		fmt.Printf("err: %v\n", err)
		log.Fatal()
	}
	return DB
}
