package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter atomic.Int64 // Use atomic.Int64 for an atomic counter

	var wg sync.WaitGroup
	numGoroutines := 1000

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			// Atomically increment the counter
			counter.Add(1) 
		}()
	}

	wg.Wait()

	// Atomically load the final value of the counter
	finalCount := counter.Load() 
	fmt.Println("Final counter value:", finalCount) // Expected: 1000
}
