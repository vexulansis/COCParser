package db

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

func GenerateCredentials(n int, DC *DBClient) error {
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
		task := &Task{
			ID:   i,
			Data: cred,
		}
		DC.TaskChan <- task
		// f := DBLoggerFields{
		// 	Source:      "GENERATOR",
		// 	Method:      "SEND",
		// 	Subject:     fmt.Sprintf("TASK#%d", task.ID),
		// 	Destination: "TASKCHANNEL",
		// }
		// DC.Logger.Print(f, 0)
	}
	return nil
}
