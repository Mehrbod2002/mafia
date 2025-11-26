package chat

import (
	"context"
	"sync"
	"time"
)

// Message holds lightweight chat content.
type Message struct {
	ID       string
	RoomID   string
	SenderID string
	Body     string
	SentAt   time.Time
}

// Room stores chat history in memory.
type Room struct {
	mu       sync.RWMutex
	messages []Message
}

// Hub manages rooms and message fanout.
type Hub struct {
	mu    sync.RWMutex
	rooms map[string]*Room
}

// NewHub creates a chat hub.
func NewHub() *Hub {
	return &Hub{rooms: make(map[string]*Room)}
}

// PostMessage appends a message to the room and returns it.
func (h *Hub) PostMessage(_ context.Context, msg Message) Message {
	h.mu.Lock()
	room := h.rooms[msg.RoomID]
	if room == nil {
		room = &Room{}
		h.rooms[msg.RoomID] = room
	}
	h.mu.Unlock()

	room.mu.Lock()
	defer room.mu.Unlock()
	if msg.SentAt.IsZero() {
		msg.SentAt = time.Now()
	}
	room.messages = append(room.messages, msg)
	return msg
}

// History returns the messages for the room.
func (h *Hub) History(roomID string) []Message {
	h.mu.RLock()
	room := h.rooms[roomID]
	h.mu.RUnlock()
	if room == nil {
		return nil
	}
	room.mu.RLock()
	defer room.mu.RUnlock()
	return append([]Message(nil), room.messages...)
}
