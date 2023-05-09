package athena

import (
	"sync"
	"time"
)

type TokenLimiter struct {
	maxsize    uint64
	tokens     uint64
	lastAccess time.Time
	every      time.Duration
	mu         sync.Mutex
}

func NewTokenLimiter(maxsize uint64, every time.Duration) *TokenLimiter {
	return &TokenLimiter{
		maxsize:    maxsize,
		tokens:     maxsize,
		every:      every,
		lastAccess: time.Now(),
	}
}

func (t *TokenLimiter) withinRefreshWindow() bool {
	nw := time.Now()

	if nw.After(t.lastAccess.Add(t.every)) {
		return false
	}

	return true
}

func (t *TokenLimiter) Get() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	if !t.withinRefreshWindow() {
		t.tokens = t.maxsize
	}

	if t.tokens > 0 {
		t.tokens--
		return true
	}

	return false
}
