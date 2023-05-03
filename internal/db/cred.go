package db

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Credentials struct {
	Email    string
	Password string
}

func GenerateCredentials(n int, wp *WorkerPool) error {
	credExample := &Credentials{}
	_, err := toml.DecodeFile("cred.toml", &credExample)
	if err != nil {
		return err
	}
	for i := 1; i <= n; i++ {
		cred := Credentials{
			Email:    fmt.Sprintf(credExample.Email, i),
			Password: credExample.Password,
		}
		task := &InsertCredTask{
			Credentials: cred,
			WG:          wp.WG,
		}
		wp.Tasks <- task
		f := DBLoggerFields{
			Source:  "GENERATOR",
			Method:  "SEND",
			Subject: "TASK",
		}
		wp.DC.Logger.Print(f, 0)
	}
	return nil
}
