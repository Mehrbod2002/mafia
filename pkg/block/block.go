package block

import "sync"

// List tracks blocked relationships between users.
type List struct {
	mu      sync.RWMutex
	blocked map[string]map[string]struct{}
}

// NewList constructs a new block list.
func NewList() *List {
	return &List{blocked: make(map[string]map[string]struct{})}
}

// Block adds target to the caller's block set.
func (l *List) Block(userID, targetID string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.blocked[userID] == nil {
		l.blocked[userID] = make(map[string]struct{})
	}
	l.blocked[userID][targetID] = struct{}{}
}

// Unblock removes target from the caller's block set.
func (l *List) Unblock(userID, targetID string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.blocked[userID] != nil {
		delete(l.blocked[userID], targetID)
	}
}

// IsBlocked reports whether the target is blocked by the caller.
func (l *List) IsBlocked(userID, targetID string) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	_, ok := l.blocked[userID][targetID]
	return ok
}
