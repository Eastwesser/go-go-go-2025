package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	ctx := context.TODO()

	// Define the maximum number of concurrent workers
	maxWorkers := runtime.GOMAXPROCS(0) // Use the number of available CPU cores
	if maxWorkers == 0 {
		maxWorkers = 4 // Fallback if GOMAXPROCS returns 0
	}

	// Create a new weighted semaphore with a capacity of maxWorkers
	sem := semaphore.NewWeighted(int64(maxWorkers))

	var wg sync.WaitGroup // Used to wait for all goroutines to finish

	// Simulate processing a list of tasks
	tasks := make([]int, 10)
	for i := range tasks {
		tasks[i] = i + 1
	}

	fmt.Printf("Starting processing with max %d concurrent workers...\n", maxWorkers)

	for _, taskID := range tasks {
		// Acquire a semaphore weight of 1, blocking if maxWorkers are already active
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore for task %d: %v", taskID, err)
			break
		}

		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			defer sem.Release(1) // Release the semaphore weight when the goroutine finishes

			fmt.Printf("Worker processing task %d...\n", id)
			time.Sleep(time.Duration(id) * 100 * time.Millisecond) // Simulate work
			fmt.Printf("Worker finished task %d.\n", id)
		}(taskID)
	}

	wg.Wait() // Wait for all goroutines to complete
	fmt.Println("All tasks completed.")
}
