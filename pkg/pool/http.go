package pool

import "errors"

func (w *Worker) Request(task Message) error {
	switch task.Type {
	case "LOGIN":
	case "CREATEKEY":
	case "REVOKEKEY":
	case "GETKEYS":
	default:
		return errors.New("Incorrect task type")
	}
	return nil
}
