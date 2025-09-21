package main

import "fmt"

func callCallbacks(a, b func()) {
	go a()
	go b()
}

func main() {
	example()
}

func example() {
	firstDone := make(chan struct{})
	secondDone := make(chan struct{})

	callCallbacks(
		func() {
			fmt.Printf("a")
			close(firstDone)
			// firstDone = nil // после использования делаем канал ниловым
		},
		func() {
			fmt.Printf("b")
			close(secondDone)
			// secondDone = nil // после использования делаем канал ниловым
		},
	)

	count := 0
	for count < 2 {
		select {
			case <-firstDone:  // Блокируется если firstDone == nil
				count++
				// firstDone = nil // доп. защита
			case <-secondDone:  // Блокируется если secondDone == nil
				count++
				// secondDone = nil // доп. защита
		}
	}
	fmt.Printf("%d\n", count)
}
