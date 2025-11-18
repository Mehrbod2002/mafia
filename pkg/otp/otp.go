package otp

import (
	"crypto/rand"
	"sync"
	"time"
)

// Code represents a one-time password with expiry.
type Code struct {
	Value     string
	ExpiresAt time.Time
}

// Manager issues and validates OTPs stored in memory.
type Manager struct {
	mu    sync.Mutex
	codes map[string]Code
}

// NewManager constructs a Manager.
func NewManager() *Manager {
	return &Manager{codes: make(map[string]Code)}
}

// Generate creates an OTP for the provided key and validity period.
func (m *Manager) Generate(key string, ttl time.Duration) Code {
	code := Code{
		Value:     randomDigits(6),
		ExpiresAt: time.Now().Add(ttl),
	}
	m.mu.Lock()
	m.codes[key] = code
	m.mu.Unlock()
	return code
}

// Validate checks that the provided code matches and is not expired.
func (m *Manager) Validate(key, value string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	code, ok := m.codes[key]
	if !ok || time.Now().After(code.ExpiresAt) {
		return false
	}
	if code.Value == value {
		delete(m.codes, key)
		return true
	}
	return false
}

func randomDigits(length int) string {
	b := make([]byte, length)
	_, _ = rand.Read(b)
	for i := range b {
		b[i] = byte('0' + int(b[i])%10)
	}
	return string(b)
}
