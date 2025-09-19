package main

import (
	"fmt"
	"sync"
	"time"
)

func gen(data []int, delay time.Duration) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range data {
			time.Sleep(delay)
			out <- v
		}
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	res := make(chan int)

	// Функция для чтения из одного канала и записи в результат
	output := func(ch <-chan int) {
		defer wg.Done()
		for v := range ch {
			res <- v
		}
	}

	wg.Add(len(cs))
	
	// Запускаем горутину для каждого входного канала
	for _, ch := range cs {
		go output(ch)
	}

	// Закрываем результирующий канал когда все горутины завершатся
	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

func main() {
	c1 := gen([]int{1, 2, 3}, 100*time.Millisecond)
	c2 := gen([]int{4, 5, 6}, 150*time.Millisecond)
	c3 := gen([]int{7, 8, 9}, 50*time.Millisecond)
	
	merged := merge(c1, c2, c3)
	
	// Читаем все значения из объединенного канала
	for v := range merged {
		fmt.Println(v)
	}
}
