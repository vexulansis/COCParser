package pool

import (
	"fmt"
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
	Name       string
	Time       Time
	Received   atomic.Uint64
	Processed  atomic.Uint64
	Speed      float64
	Efficiency float64
}

// New PoolManager example
func NewPoolManager(name string) *PoolManager {
	return &PoolManager{
		Name: name,
	}
}

// Calculates operation time, speed and efficiency
func (m *PoolManager) Benchmark() {
	m.Time.Operated = m.Time.End.Sub(m.Time.Begin)
	m.Speed = float64(m.Processed.Load()) / m.Time.Operated.Seconds()
	m.Efficiency = float64(m.Processed.Load()) / float64(m.Received.Load()) * 100
}

// Format printing stats
func (m *PoolManager) Stats() {
	m.Benchmark()
	fmt.Printf("\nPool %s\n"+
		"Started on: %s\n"+
		"Ended on: %s\n"+
		"Time operated: %f s\n"+
		"Received: %d tasks\n"+
		"Processed: %d tasks\n"+
		"With average speed: %.0f t/s\n"+
		"With efficiency: %.1f %%\n",
		m.Name,
		m.Time.Begin.Format(timeformat),
		m.Time.End.Format(timeformat),
		m.Time.Operated.Seconds(),
		m.Received.Load(),
		m.Processed.Load(),
		m.Speed,
		m.Efficiency)
}

// Benchmarking ErrorHandler
type ErrorManager struct {
	Name     string
	Time     Time
	Received atomic.Uint64
}

// New ErrorManager example
func NewErrorManager(name string) *ErrorManager {
	return &ErrorManager{
		Name: name,
	}
}

// Calculates operation time
func (m *ErrorManager) Benchmark() {
	m.Time.Operated = m.Time.End.Sub(m.Time.Begin)
}

// Format printing stats
func (m *ErrorManager) Stats() {
	m.Benchmark()
	fmt.Printf("\nErrorHandler %s\n"+
		"Started on: %s\n"+
		"Ended on: %s\n"+
		"Time operated: %f s\n"+
		"Received: %d errors\n",
		m.Name,
		m.Time.Begin.Format(timeformat),
		m.Time.End.Format(timeformat),
		m.Time.Operated.Seconds(),
		m.Received.Load())
}

// Benchmarking Generator
type GeneratorManager struct {
	Name      string
	Time      Time
	Generated atomic.Uint64
	Speed     float64
}

// New GeneratorManager example
func NewGeneratorManager(name string) *GeneratorManager {
	return &GeneratorManager{
		Name: name,
	}
}

// Calculates operation time and speed
func (m *GeneratorManager) Benchmark() {
	m.Time.Operated = m.Time.End.Sub(m.Time.Begin)
	m.Speed = float64(m.Generated.Load()) / m.Time.Operated.Seconds()
}

// Format printing stats
func (m *GeneratorManager) Stats() {
	m.Benchmark()
	fmt.Printf("\nGenerator %s\n"+
		"Started on: %s\n"+
		"Ended on: %s\n"+
		"Time operated: %f s\n"+
		"Generated: %d tasks\n"+
		"With average speed: %.0f t/s\n",
		m.Name,
		m.Time.Begin.Format(timeformat),
		m.Time.End.Format(timeformat),
		m.Time.Operated.Seconds(),
		m.Generated.Load(),
		m.Speed)
}
