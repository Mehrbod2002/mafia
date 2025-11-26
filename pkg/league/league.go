package league

import "time"

// Season captures metadata about a competition window.
type Season struct {
	ID       string
	Name     string
	StartsAt time.Time
	EndsAt   time.Time
}

// Record represents a user's performance in a season.
type Record struct {
	UserID string
	Points int
}

// Table stores records keyed by season.
type Table struct {
	seasons map[string]Season
	records map[string][]Record
}

// NewTable constructs a seasonal leaderboard.
func NewTable() *Table {
	return &Table{
		seasons: make(map[string]Season),
		records: make(map[string][]Record),
	}
}

// UpsertSeason registers or updates a season definition.
func (t *Table) UpsertSeason(season Season) {
	t.seasons[season.ID] = season
}

// AddRecord adds a score entry for a season.
func (t *Table) AddRecord(seasonID string, record Record) {
	t.records[seasonID] = append(t.records[seasonID], record)
}

// SeasonRecords returns the records for a given season.
func (t *Table) SeasonRecords(seasonID string) []Record {
	return append([]Record(nil), t.records[seasonID]...)
}
