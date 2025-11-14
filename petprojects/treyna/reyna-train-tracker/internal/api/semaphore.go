package api

import (
	"context"
)

// Semaphore ограничивает количество одновременных операций
// Паттерн: Semaphore для контроля concurrent access
type Semaphore struct {
	sem chan struct{}
}

// NewSemaphore создаёт новый семафор с заданным лимитом
func NewSemaphore(maxConcurrent int) *Semaphore {
	return &Semaphore{
		sem: make(chan struct{}, maxConcurrent),
	}
}

// Acquire захватывает семафор (блокируется если достигнут лимит)
func (s *Semaphore) Acquire() {
	s.sem <- struct{}{}
}

// Release освобождает семафор
func (s *Semaphore) Release() {
	<-s.sem
}

// AcquireContext захватывает семафор с учётом контекста
func (s *Semaphore) AcquireContext(ctx context.Context) error {
	select {
	case s.sem <- struct{}{}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// TryAcquire пытается захватить семафор без блокировки
func (s *Semaphore) TryAcquire() bool {
	select {
	case s.sem <- struct{}{}:
		return true
	default:
		return false
	}
}

