// internal/cache/enhanced_cache.go
package cache

import (
	"fmt"
	"sync/atomic"
	"time"
)

type EnhancedCache struct {
    cache     *InMemoryCache[interface{}]
    hits      atomic.Uint64
    misses    atomic.Uint64
    evictions atomic.Uint64
}

func NewEnhancedCache() *EnhancedCache {
    return &EnhancedCache{
        cache: NewInMemoryCache[interface{}](),
    }
}

func (ec *EnhancedCache) Get(key string) (interface{}, bool) {
    value, ok := ec.cache.Get(key)
    if ok {
        ec.hits.Add(1)
    } else {
        ec.misses.Add(1)
    }
    return value, ok
}

func (ec *EnhancedCache) Set(key string, value interface{}, ttl time.Duration) {
    ec.cache.Set(key, value, ttl)
}

func (ec *EnhancedCache) GetStats() map[string]interface{} {
    hitRate := float64(0)
    total := ec.hits.Load() + ec.misses.Load()
    if total > 0 {
        hitRate = float64(ec.hits.Load()) / float64(total) * 100
    }
    
    return map[string]interface{}{
        "hits":       ec.hits.Load(),
        "misses":     ec.misses.Load(),
        "hit_rate":   fmt.Sprintf("%.1f%%", hitRate),
        "size":       ec.cache.Size(),
        "evictions":  ec.evictions.Load(),
    }
}
