package api

import (
	"sync/atomic"
)

// Worker представляет обработчика запросов
type Worker struct {
	ID       int
	Load     atomic.Uint64 // Текущая нагрузка на воркера
	IsActive bool
}

// LoadBalancer распределяет нагрузку между воркерами
// Паттерн: Round-robin Load Balancer с отслеживанием нагрузки
type LoadBalancer struct {
	workers []*Worker
	next    atomic.Uint64 // Индекс следующего воркера (для round-robin)
}

// NewLoadBalancer создаёт новый load balancer
func NewLoadBalancer(numWorkers int) *LoadBalancer {
	lb := &LoadBalancer{
		workers: make([]*Worker, numWorkers),
	}

	for i := 0; i < numWorkers; i++ {
		lb.workers[i] = &Worker{
			ID:       i + 1,
			IsActive: true,
		}
	}

	return lb
}

// GetNextWorker возвращает следующего воркера (round-robin)
func (lb *LoadBalancer) GetNextWorker() *Worker {
	n := lb.next.Add(1)
	index := int((n - 1) % uint64(len(lb.workers)))
	worker := lb.workers[index]
	worker.Load.Add(1)
	return worker
}

// GetLeastLoadedWorker возвращает воркера с наименьшей нагрузкой
func (lb *LoadBalancer) GetLeastLoadedWorker() *Worker {
	var leastLoaded *Worker
	minLoad := ^uint64(0) // Максимальное значение uint64

	for _, worker := range lb.workers {
		if worker.IsActive {
			load := worker.Load.Load()
			if load < minLoad {
				minLoad = load
				leastLoaded = worker
			}
		}
	}

	if leastLoaded != nil {
		leastLoaded.Load.Add(1)
	}

	return leastLoaded
}

// ReleaseWorker освобождает воркера после выполнения задачи
func (lb *LoadBalancer) ReleaseWorker(worker *Worker) {
	if worker.Load.Load() > 0 {
		worker.Load.Add(^uint64(0)) // Декремент (Add(-1))
	}
}

// GetWorkerStats возвращает статистику по воркерам
func (lb *LoadBalancer) GetWorkerStats() []map[string]interface{} {
	stats := make([]map[string]interface{}, len(lb.workers))

	for i, worker := range lb.workers {
		stats[i] = map[string]interface{}{
			"id":     worker.ID,
			"load":   worker.Load.Load(),
			"active": worker.IsActive,
		}
	}

	return stats
}

