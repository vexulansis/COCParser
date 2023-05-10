package generator

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

var timeformat = "15:04:05"
var timeDur time.Duration

type GeneratorManager struct {
	Generator *Generator
	Time      Time
	Delay     time.Duration
	Generated atomic.Uint64
	Speed     float64
}
type Time struct {
	Begin    time.Time
	End      time.Time
	Operated time.Duration
}

func NewGeneratorManager(generator *Generator) *GeneratorManager {
	return &GeneratorManager{
		Generator: generator,
		Delay:     time.Duration(generator.Config.Delay) * time.Second,
	}
}

// Starts printing stats with delay
func (m *GeneratorManager) Start(ctx context.Context) {
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
func (m *GeneratorManager) Print() {
	m.CollectData()
	fmt.Printf("m.Time.Operated: %v\n", m.Time.Operated)
	fmt.Printf("m.Generated: %v\n", m.Generated)
	fmt.Printf("m.Speed: %.0f\n", m.Speed)
}

// Collecting data
func (m *GeneratorManager) CollectData() {
	m.Time.Operated = time.Now().Sub(m.Time.Begin)
	m.Speed = float64(m.Generated.Load()) / m.Time.Operated.Seconds()
}
