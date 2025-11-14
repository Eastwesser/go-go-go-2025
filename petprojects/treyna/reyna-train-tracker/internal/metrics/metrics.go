package metrics

import (
	"fmt"
	"sync/atomic"
	"time"
)

type MetricsCollector struct {
	requestsProcessed atomic.Uint64
	requestDuration   atomic.Uint64 // в наносекундах
	errorsCount       atomic.Uint64
	cacheHits         atomic.Uint64
	cacheMisses       atomic.Uint64
}

func NewMetricsCollector() *MetricsCollector {
	return &MetricsCollector{}
}

func (mc *MetricsCollector) RecordRequest(duration time.Duration, success bool) {
	mc.requestsProcessed.Add(1)
	mc.requestDuration.Add(uint64(duration.Nanoseconds()))
	if !success {
		mc.errorsCount.Add(1)
	}
}

func (mc *MetricsCollector) RecordCacheHit() {
	mc.cacheHits.Add(1)
}

func (mc *MetricsCollector) RecordCacheMiss() {
	mc.cacheMisses.Add(1)
}

func (mc *MetricsCollector) GetMetrics() map[string]interface{} {
	totalRequests := mc.requestsProcessed.Load()
	
	// Calculate average duration
	avgDuration := time.Duration(0)
	if totalRequests > 0 {
		avgNs := mc.requestDuration.Load() / totalRequests
		avgDuration = time.Duration(avgNs)
	}
	
	// Calculate error rate
	errorRate := 0.0
	if totalRequests > 0 {
		errorRate = float64(mc.errorsCount.Load()) / float64(totalRequests) * 100
	}
	
	// Calculate cache hit rate
	totalCacheRequests := mc.cacheHits.Load() + mc.cacheMisses.Load()
	cacheHitRate := 0.0
	if totalCacheRequests > 0 {
		cacheHitRate = float64(mc.cacheHits.Load()) / float64(totalCacheRequests) * 100
	}
	
	return map[string]interface{}{
		"total_requests":        totalRequests,
		"error_rate_percent":    fmt.Sprintf("%.2f%%", errorRate),
		"avg_request_time":      avgDuration.String(),
		"total_errors":          mc.errorsCount.Load(),
		"cache_hits":            mc.cacheHits.Load(),
		"cache_misses":          mc.cacheMisses.Load(),
		"cache_hit_rate":        fmt.Sprintf("%.1f%%", cacheHitRate),
		"total_cache_requests":  totalCacheRequests,
	}
}

func (mc *MetricsCollector) Reset() {
	mc.requestsProcessed.Store(0)
	mc.requestDuration.Store(0)
	mc.errorsCount.Store(0)
	mc.cacheHits.Store(0)
	mc.cacheMisses.Store(0)
}