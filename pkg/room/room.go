package room

import "sync"

// Room models a multiplayer lobby.
type Room struct {
	ID      string
	HostID  string
	Members []string
	Locked  bool
}

// Manager manages rooms and membership.
type Manager struct {
	mu    sync.RWMutex
	rooms map[string]Room
}

// NewManager constructs a Manager.
func NewManager() *Manager {
	return &Manager{rooms: make(map[string]Room)}
}

// Create initializes a room with a host.
func (m *Manager) Create(room Room) Room {
	m.mu.Lock()
	defer m.mu.Unlock()
	room.Members = append(room.Members, room.HostID)
	m.rooms[room.ID] = room
	return room
}

// Join adds a user to the room if not locked.
func (m *Manager) Join(roomID, userID string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	room, ok := m.rooms[roomID]
	if !ok || room.Locked {
		return false
	}
	for _, id := range room.Members {
		if id == userID {
			m.rooms[roomID] = room
			return true
		}
	}
	room.Members = append(room.Members, userID)
	m.rooms[roomID] = room
	return true
}

// Leave removes a user from the room.
func (m *Manager) Leave(roomID, userID string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	room := m.rooms[roomID]
	filtered := room.Members[:0]
	for _, id := range room.Members {
		if id != userID {
			filtered = append(filtered, id)
		}
	}
	room.Members = filtered
	m.rooms[roomID] = room
}

// Get returns a room by id.
func (m *Manager) Get(roomID string) (Room, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	r, ok := m.rooms[roomID]
	return r, ok
}
