package main

import (
	"fmt"
)

// generator produces a stream of integers
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

// square squares each integer received from the input channel
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

// sum sums all integers received from the input channel
func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		total := 0
		for n := range in {
			total += n
		}
		out <- total
	}()
	return out
}

func main() {
	// Create a pipeline: generator -> square -> sum
	numbers := []int{1, 2, 3, 4, 5}
	
	// Stage 1: Generate numbers
	c1 := generator(numbers...)

	// Stage 2: Square numbers
	c2 := square(c1)

	// Stage 3: Sum squared numbers
	c3 := sum(c2)

	// Consume the final result
	fmt.Println("Sum of squares:", <-c3) // Expected: 55 (1*1 + 2*2 + 3*3 + 4*4 + 5*5)
}
