package voice

import "sync"

// Stream represents an active voice session.
type Stream struct {
	ID       string
	RoomID   string
	Speakers []string
}

// Router manages voice streams per room.
type Router struct {
	mu      sync.RWMutex
	streams map[string]Stream
}

// NewRouter creates a Router instance.
func NewRouter() *Router {
	return &Router{streams: make(map[string]Stream)}
}

// Start registers a stream for a room.
func (r *Router) Start(stream Stream) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.streams[stream.RoomID] = stream
}

// AddSpeaker adds a speaker to a stream if present.
func (r *Router) AddSpeaker(roomID, userID string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	stream, ok := r.streams[roomID]
	if !ok {
		return false
	}
	for _, id := range stream.Speakers {
		if id == userID {
			r.streams[roomID] = stream
			return true
		}
	}
	stream.Speakers = append(stream.Speakers, userID)
	r.streams[roomID] = stream
	return true
}

// Stop removes the stream for a room.
func (r *Router) Stop(roomID string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.streams, roomID)
}

// Stream returns the stream for a room.
func (r *Router) Stream(roomID string) (Stream, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	s, ok := r.streams[roomID]
	return s, ok
}
