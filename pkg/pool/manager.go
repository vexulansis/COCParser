package pool

import (
	"sync/atomic"
	"time"
)

var timeformat = "15:04:05"

// Time pattern for all managers
type Time struct {
	Begin    time.Time
	End      time.Time
	Operated time.Duration
}

// Benchmarking Pool
type PoolManager struct {
	Time       Time
	Received   atomic.Uint64
	Processed  atomic.Uint64
	Speed      float64
	Efficiency float64
}

// New PoolManager example
func NewPoolManager() *PoolManager {
	return &PoolManager{}
}

// Calculates operation time, speed and efficiency
func (m *PoolManager) Benchmark() {
	m.Time.Operated = m.Time.End.Sub(m.Time.Begin)
	m.Speed = float64(m.Processed.Load()) / m.Time.Operated.Seconds()
	m.Efficiency = float64(m.Received.Load()) / float64(m.Processed.Load())
}

// Benchmarking ErrorHandler
type ErrorManager struct {
	Time     Time
	Received atomic.Uint64
}

// New ErrorManager example
func NewErrorManager() *ErrorManager {
	return &ErrorManager{}
}

// Calculates operation time
func (m *ErrorManager) Benchmark() {
	m.Time.Operated = m.Time.End.Sub(m.Time.Begin)
}

// Benchmarking Generator
type GeneratorManager struct {
	Time      Time
	Processed atomic.Uint64
	Speed     float64
}

// New GeneratorManager example
func NewGeneratorManager() *GeneratorManager {
	return &GeneratorManager{}
}

// Calculates operation time and speed
func (m *GeneratorManager) Benchmark() {
	m.Time.Operated = m.Time.End.Sub(m.Time.Begin)
	m.Speed = float64(m.Processed.Load()) / m.Time.Operated.Seconds()
}
