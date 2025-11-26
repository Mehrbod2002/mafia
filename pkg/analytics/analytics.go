package analytics

import "sync"

// Tracker collects lightweight analytics counters for events.
type Tracker struct {
	mu       sync.RWMutex
	counters map[string]int
}

// NewTracker creates an empty tracker.
func NewTracker() *Tracker {
	return &Tracker{counters: make(map[string]int)}
}

// Count increments the counter for the given event key.
func (t *Tracker) Count(event string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.counters[event]++
}

// Snapshot returns a copy of current counters for reporting.
func (t *Tracker) Snapshot() map[string]int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	snapshot := make(map[string]int, len(t.counters))
	for k, v := range t.counters {
		snapshot[k] = v
	}
	return snapshot
}
