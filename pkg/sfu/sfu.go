package sfu

import "sync"

// Subscriber is notified when media packets are available.
type Subscriber interface {
	Handle(data []byte)
}

// Node represents a simple selective forwarding unit dispatcher.
type Node struct {
	mu          sync.RWMutex
	subscribers map[string][]Subscriber
}

// NewNode constructs a new Node.
func NewNode() *Node {
	return &Node{subscribers: make(map[string][]Subscriber)}
}

// Subscribe registers a subscriber for a room.
func (n *Node) Subscribe(roomID string, s Subscriber) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.subscribers[roomID] = append(n.subscribers[roomID], s)
}

// Publish delivers a media buffer to all subscribers of a room.
func (n *Node) Publish(roomID string, payload []byte) {
	n.mu.RLock()
	subs := append([]Subscriber(nil), n.subscribers[roomID]...)
	n.mu.RUnlock()
	for _, s := range subs {
		s.Handle(payload)
	}
}
