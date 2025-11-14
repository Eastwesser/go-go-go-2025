package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"reyna-train-tracker/internal/api"
	"reyna-train-tracker/internal/config"
	"reyna-train-tracker/internal/metrics"
	"reyna-train-tracker/internal/models"
	"reyna-train-tracker/internal/tracker"
	"reyna-train-tracker/internal/utils"
)

func main() {
	fmt.Println("🚂 ТРЕКЕР РЭЙНЫ - Система отслеживания поезда Москва-Хабаровск")
	fmt.Println(strings.Repeat("=", 80))

	// Загружаем конфигурацию из environment variables
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("❌ Ошибка загрузки конфигурации: %v", err)
	}

	fmt.Printf("⚙️  Конфигурация загружена:\n")
	fmt.Printf("   Максимум одновременных запросов: %d\n", cfg.MaxConcurrentRequests)
	fmt.Printf("   Лимит запросов в секунду: %d\n", cfg.RateLimitPerSecond)
	fmt.Printf("   Время жизни кэша: %v\n", cfg.CacheTTL)
	fmt.Printf("   Количество воркеров: %d\n", cfg.NumWorkers)
	fmt.Printf("   Максимум повторов: %d\n", cfg.MaxRetries)
	fmt.Println()

	// Инициализируем сборщик метрик
	metricsCollector := metrics.NewMetricsCollector()

	// Загружаем данные маршрута
	trainTracker, err := tracker.NewTrainTracker(cfg.JSONDataPath)
	if err != nil {
		log.Fatalf("❌ Ошибка загрузки расписания: %v", err)
	}

	fmt.Printf("✅ Загружено станций: %d\n", len(trainTracker.Stations))
	fmt.Printf("📏 Общая дистанция: %d км\n", trainTracker.RouteData.TotalDistance)
	fmt.Printf("🕐 Начало путешествия: %s\n\n", trainTracker.RouteData.StartTime.Format("15:04 02.01.2006"))

	// Создаём обработчик вопросов с конфигурацией и метриками
	handler := api.NewQuestionHandlerWithConfig(trainTracker, cfg, metricsCollector)

	// Текущее время (можно изменить для тестирования)
	// Используем текущее время
	// currentTime := time.Now()
	
	// Или можно использовать конкретное время для тестирования:
	moscowTZ, _ := time.LoadLocation("Europe/Moscow")
	currentTime := time.Date(2025, 10, 11, 10, 0, 0, 0, moscowTZ)

	fmt.Printf("🕐 Текущее время: %s (Москва)\n", currentTime.Format("15:04 02.01.2006"))
	fmt.Println(strings.Repeat("=", 80))

	// Отладочная информация
	if cfg.DebugMode {
		tracker.DebugFindCurrentPosition(trainTracker.Stations, currentTime)
		trainTracker.DebugAllStations()
	}

	// Получаем текущую позицию с измерением времени
	startPos := time.Now()
	position := trainTracker.GetCurrentPosition(currentTime)
	posDuration := time.Since(startPos)
	metricsCollector.RecordRequest(posDuration, position != nil)

	if position != nil {
		fmt.Println("\n📍 ТЕКУЩАЯ ПОЗИЦИЯ:")
		fmt.Println(strings.Repeat("-", 80))
		
		if position.IsAtStation && position.CurrentStation != nil {
			fmt.Printf("🚉 Станция: %s\n", position.CurrentStation.Name)
			fmt.Printf("📏 Расстояние от Москвы: %d км\n", position.CurrentStation.DistanceFromStart)
			fmt.Printf("🌍 Часовой пояс: %s\n", position.Timezone)
			
			localTime, _ := utils.ConvertToTimezone(currentTime, position.Timezone)
			fmt.Printf("🕐 Локальное время: %s\n", localTime.Format("15:04 02.01.2006"))
		} else {
			fmt.Printf("🚂 В пути между станциями:\n")
			if position.PreviousStation != nil {
				fmt.Printf("   ├─ Предыдущая: %s\n", position.PreviousStation.Name)
			}
			if position.NextStation != nil {
				fmt.Printf("   └─ Следующая: %s\n", position.NextStation.Name)
			}
			fmt.Printf("📏 Приблизительное расстояние от Москвы: %.0f км\n", position.DistanceFromStart)
		}
		
		// Статус поезда
		startStatus := time.Now()
		status := trainTracker.GetTrainStatus(currentTime, position)
		statusDuration := time.Since(startStatus)
		metricsCollector.RecordRequest(statusDuration, true)

		if status.IsMoving {
			fmt.Printf("🚂 Статус: В ДВИЖЕНИИ\n")
			fmt.Printf("⏰ До следующей станции: %s\n", utils.FormatDuration(status.TimeToNext))
		} else {
			fmt.Printf("🛑 Статус: СТОИТ НА СТАНЦИИ\n")
			fmt.Printf("⏰ Осталось стоять: %s\n", utils.FormatDuration(status.RemainingStand))
		}
		
		// Информация о путешествии
		startJourney := time.Now()
		journeyInfo := trainTracker.GetJourneyInfo(currentTime)
		journeyDuration := time.Since(startJourney)
		metricsCollector.RecordRequest(journeyDuration, true)

		fmt.Printf("\n📅 ИНФОРМАЦИЯ О ПУТЕШЕСТВИИ:\n")
		fmt.Printf("   День путешествия: %d\n", journeyInfo.DayNumber)
		fmt.Printf("   Время в пути: %s\n", utils.FormatDuration(journeyInfo.TotalTimeInTrip))
	} else {
		fmt.Println("❌ Не удалось определить текущую позицию")
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("🔍 ОБРАБОТКА ВСЕХ 10 ВОПРОСОВ С ИСПОЛЬЗОВАНИЕМ ПАТТЕРНОВ КОНКУРЕНТНОСТИ...")
	fmt.Println("   (WaitGroup, Semaphore, RateLimiter, LoadBalancer, Fan-in/Fan-out)")
	fmt.Println(strings.Repeat("=", 80))

	// Обрабатываем все вопросы параллельно с использованием паттернов конкурентности
	startQuestions := time.Now()
	// Для сравнения производительности
	startOld := time.Now()
	oldResults := handler.ProcessAllQuestions(currentTime)
	oldDuration := time.Since(startOld)

	startNew := time.Now()  
	newResults := handler.ProcessAllQuestionsWithRetry(currentTime)
	newDuration := time.Since(startNew)

	// Используйте ту версию, которая лучше сработала
	var results []models.QuestionResult
	if newDuration < oldDuration * 2 { // Если новая версия не значительно медленнее
		results = newResults
		fmt.Printf("✅ Использована улучшенная версия с retry (время: %v)\n", newDuration)
	} else {
		results = oldResults  
		fmt.Printf("✅ Использована стандартная версия (время: %v)\n", oldDuration)
	}
	questionsDuration := time.Since(startQuestions)
	metricsCollector.RecordRequest(questionsDuration, len(results) == 10)

	// Сортируем результаты по номеру вопроса
	sort.Slice(results, func(i, j int) bool {
		return results[i].QuestionNumber < results[j].QuestionNumber
	})

	// Выводим результаты
	printResults(results)

	// Статистика использования
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("📊 СТАТИСТИКА ИСПОЛЬЗОВАНИЯ:")
	fmt.Println(strings.Repeat("-", 80))
	
	stats := trainTracker.GetStatistics()
	fmt.Printf("Всего запросов: %v\n", stats["total_requests"])
	fmt.Printf("Размер кэша: %v записей\n", stats["cache_size"])
	
	if questionStats, ok := stats["question_counters"].(map[string]uint64); ok {
		fmt.Println("\nЗапросов по вопросам:")
		for i := 1; i <= 10; i++ {
			key := fmt.Sprintf("question_%d", i)
			fmt.Printf("  Вопрос %d: %d раз(а)\n", i, questionStats[key])
		}
	}

	// Статистика Load Balancer
	fmt.Println("\n📊 СТАТИСТИКА LOAD BALANCER:")
	workerStats := handler.LoadBalancer.GetWorkerStats()
	for _, stat := range workerStats {
		fmt.Printf("  Worker %v: нагрузка = %v, активен = %v\n", 
			stat["id"], stat["load"], stat["active"])
	}

	// Статистика Rate Limiter
	fmt.Printf("\n📊 RATE LIMITER: доступно токенов = %d\n", handler.RateLimiter.GetTokenCount())

	// Метрики производительности
	fmt.Println("\n📈 МЕТРИКИ ПРОИЗВОДИТЕЛЬНОСТИ:")
	performanceMetrics := metricsCollector.GetMetrics()
	fmt.Printf("  Всего обработано запросов: %v\n", performanceMetrics["total_requests"])
	fmt.Printf("  Среднее время запроса: %v\n", performanceMetrics["avg_request_time"])
	fmt.Printf("  Процент ошибок: %v\n", performanceMetrics["error_rate_percent"])
	fmt.Printf("  Попаданий в кэш: %v\n", performanceMetrics["cache_hits"])
	fmt.Printf("  Промахов кэша: %v\n", performanceMetrics["cache_misses"])
	fmt.Printf("  Эффективность кэша: %v\n", performanceMetrics["cache_hit_rate"])
	fmt.Printf("  Время обработки 10 вопросов: %v\n", questionsDuration)

	// Время выполнения программы
	totalDuration := time.Since(startPos)
	fmt.Printf("\n⏱️  Общее время выполнения программы: %v\n", totalDuration)

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("✅ Программа успешно завершена!")
	fmt.Println(strings.Repeat("=", 80))
}

// printResults красиво выводит результаты всех вопросов
func printResults(results []models.QuestionResult) {
	emojis := []string{"", "🕐", "🏁", "🚂", "📅", "📏", "⏰", "🌍", "💬", "💬", "🗺️"}
	
	for _, result := range results {
		emoji := ""
		if result.QuestionNumber > 0 && result.QuestionNumber < len(emojis) {
			emoji = emojis[result.QuestionNumber]
		}
		
		fmt.Printf("\n%s %d️⃣  %s\n", emoji, result.QuestionNumber, result.QuestionText)
		fmt.Println("   " + strings.Repeat("-", 76))
		
		if answerMap, ok := result.Answer.(map[string]interface{}); ok {
			// Специальная обработка для каждого вопроса
			switch result.QuestionNumber {
			case 1: // Локальное время
				fmt.Printf("   🕐 Локальное время: %v\n", answerMap["local_time"])
				fmt.Printf("   🌍 Часовой пояс: %v\n", answerMap["timezone"])
				
			case 2: // Текущая станция
				if answerMap["at_station"] == true {
					fmt.Printf("   🚉 Станция: %v\n", answerMap["station"])
					fmt.Printf("   📏 Расстояние от Москвы: %v км\n", answerMap["distance_from_moscow"])
				} else {
					fmt.Printf("   🚂 Между станциями:\n")
					fmt.Printf("      Предыдущая: %v\n", answerMap["previous"])
					fmt.Printf("      Следующая: %v\n", answerMap["next"])
					fmt.Printf("   📏 Расстояние от Москвы: ~%v км\n", answerMap["distance_from_moscow"])
				}
				
			case 3: // Статус поезда
				status := answerMap["status"]
				fmt.Printf("   %s\n", status)
				if status == "СТОИТ" {
					fmt.Printf("   🚉 Станция: %v\n", answerMap["station"])
					fmt.Printf("   ⏰ Время стоянки: %v\n", answerMap["stand_duration"])
					fmt.Printf("   ⏳ Осталось стоять: %v\n", answerMap["remaining_stand"])
				} else {
					fmt.Printf("   📍 От: %v\n", answerMap["from"])
					fmt.Printf("   📍 До: %v\n", answerMap["to"])
					fmt.Printf("   ⏰ Время до следующей станции: %v\n", answerMap["time_to_next"])
				}
				
			case 4: // День путешествия
				fmt.Printf("   📅 День путешествия: %v\n", answerMap["day_number"])
				fmt.Printf("   🚀 Начало: %v\n", answerMap["start_date"])
				fmt.Printf("   ⏱️  Время в пути: %v\n", answerMap["time_in_trip"])
				
			case 5: // Расстояние
				fmt.Printf("   📏 Расстояние от Москвы: %v км\n", answerMap["distance_km"])
				fmt.Printf("   📍 Местоположение: %v\n", answerMap["location"])
				
			case 6: // Следующая станция
				fmt.Printf("   🚉 Следующая станция: %v\n", answerMap["next_station"])
				fmt.Printf("   ⏰ Время прибытия: %v\n", answerMap["arrival_time"])
				fmt.Printf("   ⏳ Осталось в пути: %v\n", answerMap["time_remaining"])
				
			case 7: // Разница во времени
				fmt.Printf("   🕐 Время в Москве: %v\n", answerMap["moscow_time"])
				fmt.Printf("   🕐 Локальное время: %v\n", answerMap["local_time"])
				fmt.Printf("   ⏰ Разница: %v\n", answerMap["difference"])
				fmt.Printf("   ➡️  %v\n", answerMap["direction"])
				
			case 8: // Сообщение ей
				fmt.Printf("   📱 Время отправки (Москва): %v\n", answerMap["send_time_moscow"])
				fmt.Printf("   📨 Время получения (у неё): %v\n", answerMap["receive_time_local"])
				fmt.Printf("   ⚡ %v\n", answerMap["note"])
				
			case 9: // Сообщение от неё
				fmt.Printf("   📱 Время отправки (у неё): %v\n", answerMap["send_time_local"])
				fmt.Printf("   📨 Время получения (Москва): %v\n", answerMap["receive_time_moscow"])
				fmt.Printf("   ⚡ %v\n", answerMap["note"])
				
			case 10: // Основные станции впереди
				if stations, ok := answerMap["upcoming_stations"].([]map[string]interface{}); ok {
					fmt.Printf("   🚉 Основных станций впереди: %v\n\n", answerMap["count"])
					for i, station := range stations {
						if i >= 5 { // Выводим первые 5 станций
							fmt.Printf("   ... и ещё %d станций\n", len(stations)-5)
							break
						}
						fmt.Printf("   • %v\n", station["name"])
						fmt.Printf("     ⏰ Прибытие: %v\n", station["arrival_time"])
						fmt.Printf("     🕐 Стоянка: %v\n", station["stand_duration"])
						fmt.Printf("     📏 Расстояние: %v км\n", station["distance"])
						if i < len(stations)-1 && i < 4 {
							fmt.Println()
						}
					}
				}
			}
		}
	}
}