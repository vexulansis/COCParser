package pool

import "errors"

// Switch between possible task types
func (w *Worker) Query(task Message) error {
	switch task.Type {
	case "INSERT":
		w.QueryInsert(task.Data)
	case "UPDATE":
		w.QueryUpdate(task.Data)
	default:
		return errors.New("Incorrect task type")
	}
	return nil
}

// Switch between possible data types
func (w *Worker) QueryInsert(data any) error {
	switch data.(type) {
	case Account:
		//
	default:
		return errors.New("Incorrect data type")
	}
	return nil
}

// Switch between possible data types
func (w *Worker) QueryUpdate(data any) error {
	switch data.(type) {
	case Account:
		//
	default:
		return errors.New("Incorrect data type")
	}
	return nil
}
