package db

import "database/sql"

type DBClient struct {
	Logger    *DBLogger
	DB        *sql.DB
	ErrorChan chan error
}

func NewClient() (*DBClient, error) {
	// Creating DC example
	dbClient := &DBClient{}
	// Initializing logger
	dbClient.Logger = initDBLogger()
	// Getting DB
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	dbClient.DB = db
	// Creating error channel
	dbClient.ErrorChan = make(chan error)
	return dbClient, nil
}
