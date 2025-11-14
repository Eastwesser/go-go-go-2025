package main

import (
	"fmt"
	"testing"
)

// Бенчмарк для SET операций (только запись)
func BenchmarkCache_Set(b *testing.B) {
	cache := NewInMemoryCache()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Set(key, "value")
	}
}
/*
BenchmarkCache_Set-4     1000000              4724 ns/op             190 B/op          2 allocs/op
*/

// Бенчмарк для GET операций (только чтение)
func BenchmarkCache_Get(b *testing.B) {
	cache := NewInMemoryCache()
	// Предзаполняем кэш
	for i := 0; i < 1000; i++ {
		cache.Set(fmt.Sprintf("key%d", i), "value")
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", i%1000) // Циклически читаем
		cache.Get(key)
	}
}
/*
BenchmarkCache_Get-4       10000            281664 ns/op              29 B/op          2 allocs/op
*/

// Бенчмарк для конкурентных SET операций
func BenchmarkCache_Set_Concurrent(b *testing.B) {
	cache := NewInMemoryCache()
	
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := fmt.Sprintf("key%d", i)
			cache.Set(key, "value")
			i++
		}
	})
}
/*
BenchmarkCache_Set_Concurrent-4           729938              3366 ns/op              48 B/op       2 allocs/op
*/

// Бенчмарк для конкурентных GET операций
func BenchmarkCache_Get_Concurrent(b *testing.B) {
	cache := NewInMemoryCache()
	// Предзаполняем кэш
	for i := 0; i < 1000; i++ {
		cache.Set(fmt.Sprintf("key%d", i), "value")
	}
	
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := fmt.Sprintf("key%d", i%1000)
			cache.Get(key)
			i++
		}
	})
}
/*
BenchmarkCache_Get_Concurrent-4            10000            143718 ns/op              29 B/op       2 allocs/op
*/

// Бенчмарк для смешанной нагрузки (70% reads, 30% writes)
func BenchmarkCache_MixedWorkload(b *testing.B) {
	cache := NewInMemoryCache()
	// Предзаполняем кэш
	for i := 0; i < 1000; i++ {
		cache.Set(fmt.Sprintf("key%d", i), "value")
	}
	
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%10 < 3 { // 30% записей
				key := fmt.Sprintf("key%d", i)
				cache.Set(key, "new_value")
			} else { // 70% чтений
				key := fmt.Sprintf("key%d", i%1000)
				cache.Get(key)
			}
			i++
		}
	})
}
/*
BenchmarkCache_MixedWorkload-4             10000            153410 ns/op              25 B/op       2 allocs/op
*/

// Бенчмарк для измерения contention (высокая конкуренция за один ключ)
func BenchmarkCache_HighContention(b *testing.B) {
	cache := NewInMemoryCache()
	cache.Set("hot_key", "initial_value")
	
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// Все горутины работают с одним ключом - максимальный contention
			cache.Get("hot_key")
			cache.Set("hot_key", "updated_value")
		}
	})
}
/*
BenchmarkCache_HighContention-4             9631            382754 ns/op              16 B/op       1 allocs/op
*/

// Бенчмарк для измерения памяти
func BenchmarkCache_MemoryUsage(b *testing.B) {
	b.ReportAllocs() // Включаем отчет по аллокациям
	
	cache := NewInMemoryCache()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Set(key, "some_value_that_takes_space")
	}
}
/*
BenchmarkCache_MemoryUsage-4     1000000              5150 ns/op             190 B/op          2 allocs/op
*/

// # Все бенчмарки
// go test -bench=. -benchmempackage main

// # Только определенные бенчмарки
// go test -bench="BenchmarkCache_Set" -benchmem
// go test -bench="BenchmarkCache_Get_Concurrent" -benchmem

// # С профилированием
// go test -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out
