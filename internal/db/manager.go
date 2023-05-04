package db

import (
	"fmt"
	"sync"
	"time"
)

var timeformat = "15:04:05"

type PoolManager struct {
	Mutex          *sync.Mutex
	StartTime      time.Time
	ShutdownTime   time.Time
	Operated       time.Duration
	TasksReceived  int
	TasksProcessed int
	Speed          float64
	Efficiency     float64
}

func initPoolManager() *PoolManager {
	return &PoolManager{
		StartTime: time.Now(),
		Mutex:     &sync.Mutex{},
	}
}
func (m *PoolManager) Benchmark() {
	m.Operated = m.ShutdownTime.Sub(m.StartTime)
	m.Speed = float64(m.TasksProcessed) / float64(m.Operated.Seconds())
	m.Efficiency = float64(m.TasksProcessed) / float64(m.TasksReceived) * 100
}
func (m *PoolManager) PrintBenchmark() {
	m.Benchmark()
	format := "\nPool \n" +
		"Start time: %s\n" +
		"Shutdown time: %s\n" +
		"Time operated: %f s\n" +
		"Tasks received: %d\n" +
		"Tasks processed: %d\n" +
		"With average speed: %.0f tasks/s \n" +
		"With efficiency: %.0f %%\n"
	fmt.Printf(format, m.StartTime.Format(timeformat), m.ShutdownTime.Format(timeformat), m.Operated.Seconds(), m.TasksReceived, m.TasksProcessed, m.Speed, m.Efficiency)
}

type GeneratorManager struct {
	Mutex          *sync.Mutex
	StartTime      time.Time
	ShutdownTime   time.Time
	Operated       time.Duration
	TasksGenerated int
	Speed          float64
}

func initGeneratorManager() *GeneratorManager {
	return &GeneratorManager{
		StartTime: time.Now(),
		Mutex:     &sync.Mutex{},
	}
}
func (m *GeneratorManager) Benchmark() {
	m.Operated = m.ShutdownTime.Sub(m.StartTime)
	m.Speed = float64(m.TasksGenerated) / m.Operated.Seconds()
}
func (m *GeneratorManager) PrintBenchmark() {
	m.Benchmark()
	format := "\nGenerator \n" +
		"Start time: %s\n" +
		"Shutdown time: %s\n" +
		"Time operated: %f s\n" +
		"Tasks generated: %d\n" +
		"With average speed: %.0f tasks/s \n"
	fmt.Printf(format, m.StartTime.Format(timeformat), m.ShutdownTime.Format(timeformat), m.Operated.Seconds(), m.TasksGenerated, m.Speed)
}
