package term

import "time"

// Acceptance records when a user accepted a terms version.
type Acceptance struct {
	UserID     string
	Version    string
	AcceptedAt time.Time
}

// Registry keeps track of acceptances.
type Registry struct {
	records map[string]Acceptance
}

// NewRegistry constructs an empty Registry.
func NewRegistry() *Registry {
	return &Registry{records: make(map[string]Acceptance)}
}

// Accept stores the acceptance for the user.
func (r *Registry) Accept(userID, version string) Acceptance {
	acc := Acceptance{UserID: userID, Version: version, AcceptedAt: time.Now()}
	r.records[userID] = acc
	return acc
}

// Status returns the acceptance info for a user if present.
func (r *Registry) Status(userID string) (Acceptance, bool) {
	acc, ok := r.records[userID]
	return acc, ok
}
