package wallet

import "sync"

// Wallet tracks a balance for a user.
type Wallet struct {
	mu      sync.RWMutex
	Balance int64
}

// Credit adds funds to the wallet.
func (w *Wallet) Credit(amount int64) {
	w.mu.Lock()
	w.Balance += amount
	w.mu.Unlock()
}

// Debit removes funds if available and returns success.
func (w *Wallet) Debit(amount int64) bool {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.Balance < amount {
		return false
	}
	w.Balance -= amount
	return true
}

// BalanceOf returns the current balance.
func (w *Wallet) BalanceOf() int64 {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.Balance
}
