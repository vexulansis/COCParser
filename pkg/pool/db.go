package pool

import (
	"encoding/json"
	"errors"
)

// Switch between possible task types
func (w *Worker) Query(task Message) error {
	switch task.Type {
	case "INSERTACCOUNT":
		acc := &Account{}
		a, err := json.Marshal(task.Data)
		if err != nil {
			return err
		}
		json.Unmarshal(a, acc)
		_, err = w.Pool.DB.Exec(InsertAccountQuery, acc.ID, acc.Email, acc.Password)
		if err != nil {
			return err
		}
	case "UPDATE":
	default:
		return errors.New("Incorrect task type")
	}
	return nil
}
