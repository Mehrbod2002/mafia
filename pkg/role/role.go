package role

import "sync"

// Role expresses a named permission set.
type Role struct {
	Name        string
	Permissions []string
}

// Registry stores available roles.
type Registry struct {
	mu    sync.RWMutex
	roles map[string]Role
}

// NewRegistry builds a Registry.
func NewRegistry() *Registry {
	return &Registry{roles: make(map[string]Role)}
}

// Upsert registers or replaces a role definition.
func (r *Registry) Upsert(role Role) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.roles[role.Name] = role
}

// Get returns a role by name.
func (r *Registry) Get(name string) (Role, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	role, ok := r.roles[name]
	return role, ok
}
