package friend

import "sync"

// Manager keeps track of friendship relationships.
type Manager struct {
	mu      sync.RWMutex
	friends map[string]map[string]struct{}
}

// NewManager builds a Manager.
func NewManager() *Manager {
	return &Manager{friends: make(map[string]map[string]struct{})}
}

// Add registers mutual friendship between two users.
func (m *Manager) Add(a, b string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.friends[a] == nil {
		m.friends[a] = make(map[string]struct{})
	}
	if m.friends[b] == nil {
		m.friends[b] = make(map[string]struct{})
	}
	m.friends[a][b] = struct{}{}
	m.friends[b][a] = struct{}{}
}

// Remove deletes a friendship from both sides.
func (m *Manager) Remove(a, b string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.friends[a], b)
	delete(m.friends[b], a)
}

// List returns a copy of the user's friend IDs.
func (m *Manager) List(userID string) []string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	res := make([]string, 0, len(m.friends[userID]))
	for id := range m.friends[userID] {
		res = append(res, id)
	}
	return res
}
