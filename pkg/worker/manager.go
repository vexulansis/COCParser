package worker

import (
	"sync/atomic"
)

type WorkerManager struct {
	Worker    *Worker
	Received  atomic.Uint64
	Processed atomic.Uint64
}

func NewWorkerManager(worker *Worker) *WorkerManager {
	return &WorkerManager{
		Worker: worker,
	}
}
