package report

import "sync"

// Report captures a user generated report.
type Report struct {
	ID       string
	Reporter string
	Target   string
	Reason   string
}

// Repository stores reports in memory for later processing.
type Repository struct {
	mu      sync.RWMutex
	reports []Report
}

// NewRepository constructs a Repository.
func NewRepository() *Repository {
	return &Repository{}
}

// Add appends a new report.
func (r *Repository) Add(report Report) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.reports = append(r.reports, report)
}

// All returns a copy of all submitted reports.
func (r *Repository) All() []Report {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return append([]Report(nil), r.reports...)
}
