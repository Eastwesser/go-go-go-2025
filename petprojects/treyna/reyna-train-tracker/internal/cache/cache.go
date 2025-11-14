package cache

import (
	"sync"
	"time"

	"reyna-train-tracker/internal/models"
)

// Cache интерфейс для кэша с generic типом
// Демонстрирует использование generics в Go
type Cache[T any] interface {
	Set(key string, value T, ttl time.Duration)
	Get(key string) (T, bool)
	Delete(key string)
	Clear()
}

// InMemoryCache реализация кэша в памяти с RWMutex для защиты
// Паттерн: In-memory cache + RWMutex для concurrent access
type InMemoryCache[T any] struct {
	data map[string]models.CacheEntry[T]
	mu   sync.RWMutex // RWMutex позволяет множественное чтение, но эксклюзивную запись
}

// NewInMemoryCache создаёт новый кэш
func NewInMemoryCache[T any]() *InMemoryCache[T] {
	cache := &InMemoryCache[T]{
		data: make(map[string]models.CacheEntry[T]),
	}

	// Запускаем горутину для очистки устаревших записей
	go cache.cleanupExpired()

	return cache
}

// Set добавляет значение в кэш с TTL
func (c *InMemoryCache[T]) Set(key string, value T, ttl time.Duration) {
	c.mu.Lock()         // Эксклюзивная блокировка для записи
	defer c.mu.Unlock()

	c.data[key] = models.CacheEntry[T]{
		Value:     value,
		Timestamp: time.Now(),
		TTL:       ttl,
	}
}

// Get получает значение из кэша
func (c *InMemoryCache[T]) Get(key string) (T, bool) {
	c.mu.RLock()         // Разделяемая блокировка для чтения
	defer c.mu.RUnlock()

	entry, ok := c.data[key]
	if !ok {
		var zero T
		return zero, false
	}

	// Проверяем, не истёк ли TTL
	if entry.TTL > 0 && time.Since(entry.Timestamp) > entry.TTL {
		var zero T
		return zero, false
	}

	return entry.Value, true
}

// Delete удаляет значение из кэша
func (c *InMemoryCache[T]) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
}

// Clear очищает весь кэш
func (c *InMemoryCache[T]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data = make(map[string]models.CacheEntry[T])
}

// cleanupExpired периодически очищает устаревшие записи
func (c *InMemoryCache[T]) cleanupExpired() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for key, entry := range c.data {
			if entry.TTL > 0 && now.Sub(entry.Timestamp) > entry.TTL {
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}

// Size возвращает количество элементов в кэше
func (c *InMemoryCache[T]) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.data)
}

