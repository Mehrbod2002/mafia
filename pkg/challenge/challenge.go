package challenge

import "sync"

// Challenge represents a competitive objective between players.
type Challenge struct {
	ID       string
	Name     string
	Status   string
	Creator  string
	Opponent string
}

// Manager coordinates challenge lifecycle in memory.
type Manager struct {
	mu         sync.RWMutex
	challenges map[string]Challenge
}

// NewManager creates a Manager instance.
func NewManager() *Manager {
	return &Manager{challenges: make(map[string]Challenge)}
}

// Create registers a new pending challenge.
func (m *Manager) Create(c Challenge) Challenge {
	m.mu.Lock()
	defer m.mu.Unlock()
	if c.Status == "" {
		c.Status = "pending"
	}
	m.challenges[c.ID] = c
	return c
}

// UpdateStatus moves a challenge to a new state.
func (m *Manager) UpdateStatus(id, status string) (Challenge, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	c, ok := m.challenges[id]
	if !ok {
		return Challenge{}, false
	}
	c.Status = status
	m.challenges[id] = c
	return c, true
}

// Get fetches a challenge by id.
func (m *Manager) Get(id string) (Challenge, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	c, ok := m.challenges[id]
	return c, ok
}
