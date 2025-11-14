package main

import (
	"fmt"
	"time"
)

// Job represents a task to be processed by a worker.
type Job struct {
	ID int
	// Add other relevant data for your job
}

// Result represents the outcome of a processed job.
type Result struct {
	JobID int
	// Add other relevant data for your result
}

// worker function simulates a worker processing jobs.
func worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job.ID)
		time.Sleep(time.Second) // Simulate an expensive task
		fmt.Printf("Worker %d finished job %d\n", id, job.ID)
		results <- Result{JobID: job.ID}
	}
}

func main() {
	numWorkers := 3 // Number of concurrent workers
	numJobs := 10   // Total number of jobs to process

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// Start the workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs to the jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j}
	}
	close(jobs) // Close the jobs channel after all jobs are sent

	// Collect results from the results channel
	for a := 1; a <= numJobs; a++ {
		result := <-results
		fmt.Printf("Collected result for job %d\n", result.JobID)
	}
	close(results) // Close the results channel after all results are collected
}
