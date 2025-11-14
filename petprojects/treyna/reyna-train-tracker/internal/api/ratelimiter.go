package api

import (
	"sync"
	"time"
)

// RateLimiter ограничивает частоту запросов
// Паттерн: Token Bucket Rate Limiter
type RateLimiter struct {
	tokens         int
	maxTokens      int
	refillRate     time.Duration
	lastRefillTime time.Time
	mu             sync.Mutex
}

// NewRateLimiter создаёт новый rate limiter
// maxTokens - максимальное количество токенов
// refillRate - как часто добавляется новый токен
func NewRateLimiter(maxTokens int, refillRate time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens:         maxTokens,
		maxTokens:      maxTokens,
		refillRate:     refillRate,
		lastRefillTime: time.Now(),
	}

	// Запускаем горутину для пополнения токенов
	go rl.refillTokens()

	return rl
}

// Allow проверяет, можно ли выполнить запрос
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}

	return false
}

// Wait ждёт, пока не появится доступный токен
func (rl *RateLimiter) Wait() {
	for !rl.Allow() {
		time.Sleep(rl.refillRate / 10)
	}
}

// refillTokens периодически пополняет токены
func (rl *RateLimiter) refillTokens() {
	ticker := time.NewTicker(rl.refillRate)
	defer ticker.Stop()

	for range ticker.C {
		rl.mu.Lock()
		if rl.tokens < rl.maxTokens {
			rl.tokens++
		}
		rl.lastRefillTime = time.Now()
		rl.mu.Unlock()
	}
}

// GetTokenCount возвращает текущее количество доступных токенов
func (rl *RateLimiter) GetTokenCount() int {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	return rl.tokens
}

