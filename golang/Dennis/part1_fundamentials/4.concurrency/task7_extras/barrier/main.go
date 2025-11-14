package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, barrier *sync.WaitGroup) {
	defer wg.Done() // Signal that this worker is done with its overall task

	fmt.Printf("Worker %d: Starting phase 1\n", id)
	time.Sleep(time.Duration(id) * 100 * time.Millisecond) // Simulate work
	fmt.Printf("Worker %d: Finished phase 1\n", id)

	barrier.Done() // Signal that this worker has reached the barrier
	barrier.Wait() // Wait for all other workers to reach the barrier

	fmt.Printf("Worker %d: Starting phase 2\n", id)
	time.Sleep(time.Duration(id) * 50 * time.Millisecond) // Simulate more work
	fmt.Printf("Worker %d: Finished phase 2\n", id)
}

func main() {
	numWorkers := 3
	var wg sync.WaitGroup // To wait for all workers to complete their entire execution
	var barrier sync.WaitGroup // To synchronize workers at a specific point

	barrier.Add(numWorkers) // Initialize barrier for all workers to reach it

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, &barrier)
	}

	wg.Wait() // Wait for all workers to finish all phases
	fmt.Println("All workers completed.")
}

