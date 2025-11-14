package main

import (
	"fmt"
)

// evenGenerator creates a channel and starts a goroutine to send even numbers
// up to a specified maximum into that channel.
func evenGenerator(max int) <-chan int {
	out := make(chan int) // Create an unbuffered channel
	go func() {
		defer close(out) // Ensure the channel is closed when the goroutine finishes
		for i := 0; i <= max; i += 2 {
			out <- i // Send the even number to the channel
		}
	}()
	return out // Return the receive-only channel
}

func main() {
	// Consume values from the evenGenerator using a for-range loop
	for num := range evenGenerator(10) {
		fmt.Println(num)
	}
	fmt.Println("Generator finished.")
}
