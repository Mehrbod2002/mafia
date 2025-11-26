package domain

import (
	"encoding/json"
	"time"
)

// AbilityUsage tracks when a player last used a specific ability.
type AbilityUsage struct {
	Day   int    `json:"day"`
	Phase string `json:"phase"`
}

// PlayerAssignment keeps the active role/ability information for a player within a room.
type PlayerAssignment struct {
	Role          string                  `json:"role"`
	Team          string                  `json:"team"`
	Abilities     []string                `json:"abilities"`
	Alive         bool                    `json:"alive"`
	UsedAbilities map[string]AbilityUsage `json:"used_abilities"`
}

// VoteLog captures a single vote action during the day phase.
type VoteLog struct {
	Voter     uint      `json:"voter"`
	Target    uint      `json:"target"`
	Phase     string    `json:"phase"`
	Day       int       `json:"day"`
	Timestamp time.Time `json:"timestamp"`
}

// AbilityAction records a player using a specific ability.
type AbilityAction struct {
	UserID    uint      `json:"user_id"`
	Ability   string    `json:"ability"`
	TargetID  uint      `json:"target_id"`
	Phase     string    `json:"phase"`
	Day       int       `json:"day"`
	Timestamp time.Time `json:"timestamp"`
}

// GameState keeps the serialized state of an in-progress game.
type GameState struct {
	Phase       string                    `json:"phase"`
	DayCount    int                       `json:"day_count"`
	Assignments map[uint]PlayerAssignment `json:"assignments"`
	Votes       []VoteLog                 `json:"votes"`
	Abilities   []AbilityAction           `json:"abilities"`
}

// NewGameState creates an empty state for the given phase/day.
func NewGameState(phase string, day int) *GameState {
	return &GameState{
		Phase:       phase,
		DayCount:    day,
		Assignments: map[uint]PlayerAssignment{},
		Votes:       []VoteLog{},
		Abilities:   []AbilityAction{},
	}
}

// Serialize renders the game state into a string for persistence.
func (g *GameState) Serialize() (string, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ParseGameState converts the stored string version back into a structured GameState.
func ParseGameState(raw string) (*GameState, error) {
	if raw == "" {
		return NewGameState("", 0), nil
	}
	var gs GameState
	if err := json.Unmarshal([]byte(raw), &gs); err != nil {
		return nil, err
	}
	if gs.Assignments == nil {
		gs.Assignments = map[uint]PlayerAssignment{}
	}
	return &gs, nil
}
