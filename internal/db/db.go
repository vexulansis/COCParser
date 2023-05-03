package db

import (
	"database/sql"
	"fmt"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     int
}

func getDB() (*sql.DB, error) {
	// Creating DBConfig example
	dbConfig := &DBConfig{}
	// Getting DB configuration from .toml
	_, err := toml.DecodeFile("db.toml", &dbConfig)
	if err != nil {
		return nil, err
	}
	// Generating DB URI
	DBURI := getURI(dbConfig)
	// Getting DB example
	db, err := sql.Open("postgres", DBURI)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getURI(config *DBConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Name,
	)
}
