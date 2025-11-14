package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// worker получает задачи из jobs и отправляет результаты в results.
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Эмуляция работы
		fmt.Printf("Worker %d обработал задачу %d\n", id, job)
		results <- job * 2 // Возвращаем результат
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const numWorkers = 5
	const numJobs = 10

	jobs := make(chan int, numJobs)     // Канал для задач
	results := make(chan int, numJobs)  // Канал для результатов

	var wg sync.WaitGroup

	// Запускаем 5 воркеров
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Fan-Out: Отправляем задачи в канал
	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Закрываем канал, чтобы воркеры знали, что задач больше не будет

	// Ожидаем завершения всех воркеров
	go func() {
		wg.Wait()
		close(results) // Закрываем канал результатов
	}()

	// Fan-In: Собираем результаты
	for result := range results {
		fmt.Println("Результат:", result)
	}
}
