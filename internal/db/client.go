package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBClient struct {
	DB     *sql.DB
	Config DBConfig
}
type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     int
}

func getDB(config DBConfig) (*sql.DB, error) {
	DBURI := getURI(config)
	db, err := sql.Open("postgres", DBURI)
	if err != nil {
		return nil, err
	}
	return db, nil
}
func getURI(config DBConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Name,
	)
}
