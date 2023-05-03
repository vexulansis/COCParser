package db

import (
	"fmt"
	"sync"
)

type Task interface {
	Execute(*DBClient, int) error
}
type InsertCredTask struct {
	Credentials Credentials
	WG          *sync.WaitGroup
}

func (t *InsertCredTask) Execute(DC *DBClient, id int) error {
	if t.WG != nil {
		defer t.WG.Done()
	}
	res, err := DC.DB.Exec("insert into credentials(email,password) values ($1,$2)", t.Credentials.Email, t.Credentials.Password)
	if err != nil {
		DC.ErrorChan <- err
	}
	f := DBLoggerFields{
		Source:      fmt.Sprintf("WORKER#%d", id),
		Method:      "INSERT",
		Subject:     t.Credentials.Email,
		Destination: "credentials",
	}
	DC.Logger.Print(f, res)
	return nil
}
