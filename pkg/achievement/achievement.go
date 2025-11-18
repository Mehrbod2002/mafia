package achievement

import (
	"sync"
	"time"
)

// Achievement represents an earned accomplishment for a user.
type Achievement struct {
	ID       string
	UserID   string
	Name     string
	EarnedAt time.Time
	Metadata map[string]string
}

// Store keeps achievements in memory for quick lookups.
type Store struct {
	mu    sync.RWMutex
	items map[string][]Achievement
}

// NewStore creates a threadsafe achievement store.
func NewStore() *Store {
	return &Store{items: make(map[string][]Achievement)}
}

// Add records a new achievement for a user.
func (s *Store) Add(userID string, achievement Achievement) {
	s.mu.Lock()
	defer s.mu.Unlock()
	achievement.UserID = userID
	if achievement.EarnedAt.IsZero() {
		achievement.EarnedAt = time.Now()
	}
	s.items[userID] = append(s.items[userID], achievement)
}

// List returns all achievements for a user.
func (s *Store) List(userID string) []Achievement {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]Achievement(nil), s.items[userID]...)
}

// Has returns true when the user already has the achievement name.
func (s *Store) Has(userID, name string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, a := range s.items[userID] {
		if a.Name == name {
			return true
		}
	}
	return false
}
