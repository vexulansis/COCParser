package pool

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

var timeformat = "15:04:05"
var timeDur time.Duration

type PoolManager struct {
	Pool       *Pool
	Time       Time
	Delay      time.Duration
	Received   atomic.Uint64
	Processed  atomic.Uint64
	Speed      float64
	Efficiency float64
}
type Time struct {
	Begin    time.Time
	End      time.Time
	Operated time.Duration
}

func NewPoolManager(pool *Pool) *PoolManager {
	return &PoolManager{
		Pool:  pool,
		Delay: time.Duration(pool.Config.Delay) * time.Second,
	}
}

// Starts printing stats with delay
func (m *PoolManager) Start(ctx context.Context) {
	for {
		time.Sleep(m.Delay)
		m.Print()
		select {
		case <-ctx.Done():
			return
		default:

		}
	}
}

// Prints stats
func (m *PoolManager) Print() {
	m.CollectData()
	fmt.Printf("m.Time.Operated: %v\n", m.Time.Operated)
	fmt.Printf("m.Received: %v\n", m.Received)
	fmt.Printf("m.Processed: %v\n", m.Processed)
	fmt.Printf("m.Speed: %v\n", m.Speed)
}

// Collecting data from worker managers
func (m *PoolManager) CollectData() {
	m.Time.Operated = time.Now().Sub(m.Time.Begin)
	m.Received.Store(0)
	m.Processed.Store(0)
	for _, w := range m.Pool.Workers {
		m.Received.Add(w.Manager.Received.Load())
		m.Processed.Add(w.Manager.Processed.Load())
	}
	m.Speed = float64(m.Received.Load()) / m.Time.Operated.Seconds()
}
