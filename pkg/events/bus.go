package events

import (
	"context"
	"sync"
)

// Bus offers simple publish/subscribe semantics.
type Bus interface {
	Publish(ctx context.Context, topic string, payload interface{})
	Subscribe(topic string, handler func(context.Context, interface{}))
}

// SimpleBus is an in-memory implementation using an optional async dispatcher.
type SimpleBus struct {
	mu          sync.RWMutex
	subscribers map[string][]func(context.Context, interface{})
	dispatch    func(func()) error
}

// NewSimpleBus builds a bus that optionally dispatches handlers through a worker queue.
func NewSimpleBus(dispatch func(func()) error) *SimpleBus {
	return &SimpleBus{
		subscribers: make(map[string][]func(context.Context, interface{})),
		dispatch:    dispatch,
	}
}

// Publish delivers a payload to all registered handlers for the topic.
func (b *SimpleBus) Publish(ctx context.Context, topic string, payload interface{}) {
	b.mu.RLock()
	handlers := b.subscribers[topic]
	b.mu.RUnlock()

	for _, h := range handlers {
		handler := h
		if b.dispatch != nil {
			_ = b.dispatch(func() { handler(ctx, payload) })
			continue
		}
		go handler(ctx, payload)
	}
}

// Subscribe registers a handler for a topic.
func (b *SimpleBus) Subscribe(topic string, handler func(context.Context, interface{})) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.subscribers[topic] = append(b.subscribers[topic], handler)
}
