package group

import "sync"

// Group models a collection of users.
type Group struct {
	ID      string
	Name    string
	Members []string
}

// Directory manages groups and membership.
type Directory struct {
	mu     sync.RWMutex
	groups map[string]Group
}

// NewDirectory creates a Directory.
func NewDirectory() *Directory {
	return &Directory{groups: make(map[string]Group)}
}

// Save stores or updates a group definition.
func (d *Directory) Save(g Group) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.groups[g.ID] = g
}

// AddMember adds a user to the group if not already present.
func (d *Directory) AddMember(groupID, userID string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	group := d.groups[groupID]
	for _, id := range group.Members {
		if id == userID {
			d.groups[groupID] = group
			return
		}
	}
	group.Members = append(group.Members, userID)
	d.groups[groupID] = group
}

// Members returns a copy of member IDs for a group.
func (d *Directory) Members(groupID string) []string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return append([]string(nil), d.groups[groupID].Members...)
}
