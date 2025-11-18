package game

import "sync"

// State captures the progress of a lightweight turn-based game.
type State struct {
	ID        string
	Players   []string
	TurnIndex int
	Completed bool
}

// Engine drives the state transitions for demo purposes.
type Engine struct {
	mu     sync.RWMutex
	states map[string]State
}

// NewEngine constructs an Engine.
func NewEngine() *Engine {
	return &Engine{states: make(map[string]State)}
}

// Start records a new game with players.
func (e *Engine) Start(id string, players []string) State {
	e.mu.Lock()
	defer e.mu.Unlock()
	state := State{ID: id, Players: players}
	e.states[id] = state
	return state
}

// NextTurn advances the turn index when the game is active.
func (e *Engine) NextTurn(id string) (State, bool) {
	e.mu.Lock()
	defer e.mu.Unlock()
	state, ok := e.states[id]
	if !ok || state.Completed || len(state.Players) == 0 {
		return State{}, false
	}
	state.TurnIndex = (state.TurnIndex + 1) % len(state.Players)
	e.states[id] = state
	return state, true
}

// Complete marks a game as finished.
func (e *Engine) Complete(id string) (State, bool) {
	e.mu.Lock()
	defer e.mu.Unlock()
	state, ok := e.states[id]
	if !ok {
		return State{}, false
	}
	state.Completed = true
	e.states[id] = state
	return state, true
}
