package badge

import "sync"

// Badge captures cosmetic achievements a user can display.
type Badge struct {
	ID     string
	Name   string
	Icon   string
	Tier   string
	Hidden bool
}

// Locker manages badge ownership.
type Locker struct {
	mu     sync.RWMutex
	badges map[string][]Badge
}

// NewLocker constructs a badge locker.
func NewLocker() *Locker {
	return &Locker{badges: make(map[string][]Badge)}
}

// Grant gives a badge to a user if they do not already have it.
func (l *Locker) Grant(userID string, b Badge) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, owned := range l.badges[userID] {
		if owned.ID == b.ID {
			return
		}
	}
	l.badges[userID] = append(l.badges[userID], b)
}

// List returns all badges for a user.
func (l *Locker) List(userID string) []Badge {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return append([]Badge(nil), l.badges[userID]...)
}
