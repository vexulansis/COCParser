package db

import (
	"database/sql"
)

type DBClient struct {
	Logger    *DBLogger
	DB        *sql.DB
	TaskChan  chan *Task
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
	// Creating task channel
	dbClient.TaskChan = make(chan *Task)
	// Creating error channel
	dbClient.ErrorChan = make(chan error)
	return dbClient, nil
}
func (c *DBClient) GetCredentials() ([]Credentials, error) {
	credentials := []Credentials{}
	rows, err := c.DB.Query("select * from credentials")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		cred := Credentials{}
		err := rows.Scan(&cred.Email, &cred.Password)
		if err != nil {
			return nil, err
		}
		credentials = append(credentials, cred)
	}
	f := DBLoggerFields{
		Source:      "DBCLIENT",
		Method:      "SELECT",
		Subject:     "*",
		Destination: "credentials",
	}
	c.Logger.Print(f, 0)
	return credentials, nil
}
