package websocket

import "sync"

// Client represents a connected websocket consumer.
type Client struct {
	ID   string
	Send chan []byte
}

// Hub manages websocket clients and broadcasts.
type Hub struct {
	mu      sync.RWMutex
	clients map[string]*Client
}

// NewHub creates a websocket hub.
func NewHub() *Hub {
	return &Hub{clients: make(map[string]*Client)}
}

// Register adds a client to the hub.
func (h *Hub) Register(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[c.ID] = c
}

// Unregister removes a client and closes its channel.
func (h *Hub) Unregister(id string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if c, ok := h.clients[id]; ok {
		close(c.Send)
		delete(h.clients, id)
	}
}

// Broadcast sends the payload to all registered clients.
func (h *Hub) Broadcast(payload []byte) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	for _, client := range h.clients {
		select {
		case client.Send <- payload:
		default:
			// drop message when client is busy
		}
	}
}
