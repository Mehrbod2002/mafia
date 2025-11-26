package cache

import (
	"context"
	"sync"
	"time"
)

// Cache defines a minimal interface for caching values with an expiration.
type Cache interface {
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration)
	Get(ctx context.Context, key string) (interface{}, bool)
	Delete(ctx context.Context, key string)
}

// InMemoryCache is a lightweight cache with TTL support for demo purposes.
type InMemoryCache struct {
	mu    sync.RWMutex
	items map[string]cacheItem
}

type cacheItem struct {
	value     interface{}
	expiresAt time.Time
	hasExpiry bool
}

// NewInMemoryCache creates a new cache instance.
func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{items: make(map[string]cacheItem)}
}

// Set stores a value with an optional TTL.
func (c *InMemoryCache) Set(_ context.Context, key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item := cacheItem{value: value}
	if ttl > 0 {
		item.expiresAt = time.Now().Add(ttl)
		item.hasExpiry = true
	}
	c.items[key] = item
}

// Get retrieves a value if it exists and has not expired.
func (c *InMemoryCache) Get(_ context.Context, key string) (interface{}, bool) {
	c.mu.RLock()
	item, ok := c.items[key]
	c.mu.RUnlock()
	if !ok {
		return nil, false
	}
	if item.hasExpiry && time.Now().After(item.expiresAt) {
		c.Delete(context.Background(), key)
		return nil, false
	}
	return item.value, true
}

// Delete removes a value from the cache.
func (c *InMemoryCache) Delete(_ context.Context, key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}
