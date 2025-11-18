package stats

import "sync"

// Gauge represents a simple numeric value that can go up or down.
type Gauge struct {
	mu    sync.RWMutex
	value float64
}

// Inc increases the gauge by delta.
func (g *Gauge) Inc(delta float64) {
	g.mu.Lock()
	g.value += delta
	g.mu.Unlock()
}

// Dec decreases the gauge by delta.
func (g *Gauge) Dec(delta float64) {
	g.Inc(-delta)
}

// Value returns the current reading.
func (g *Gauge) Value() float64 {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.value
}
