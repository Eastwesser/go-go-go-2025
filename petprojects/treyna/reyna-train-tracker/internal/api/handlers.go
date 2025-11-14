package api

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"reyna-train-tracker/internal/config"
	"reyna-train-tracker/internal/metrics"
	"reyna-train-tracker/internal/models"
	"reyna-train-tracker/internal/tracker"
	"reyna-train-tracker/internal/utils"
)

// QuestionHandler обработчик вопросов о путешествии
type QuestionHandler struct {
	Tracker      *tracker.TrainTracker
	Config       *config.Config
	Metrics      *metrics.MetricsCollector
	Semaphore    *Semaphore
	RateLimiter  *RateLimiter
	LoadBalancer *LoadBalancer
}

// NewQuestionHandlerWithConfig создаёт новый обработчик вопросов с конфигурацией
func NewQuestionHandlerWithConfig(t *tracker.TrainTracker, cfg *config.Config, metrics *metrics.MetricsCollector) *QuestionHandler {
	return &QuestionHandler{
		Tracker:        t,
		Config:         cfg,
		Metrics:        metrics,
		Semaphore:      NewSemaphore(cfg.MaxConcurrentRequests),
		RateLimiter:    NewRateLimiter(cfg.RateLimitPerSecond, 1*time.Second),
		LoadBalancer:   NewLoadBalancer(cfg.NumWorkers),
	}
}

// ProcessAllQuestions обрабатывает все 10 вопросов параллельно
// Использует паттерны: Fan-out, Fan-in, WaitGroup
func (h *QuestionHandler) ProcessAllQuestions(currentTime time.Time) []models.QuestionResult {
	// Fan-out: запускаем обработку всех вопросов параллельно
	results := make(chan models.QuestionResult, 10)
	var wg sync.WaitGroup

	// Получаем текущую позицию один раз для всех вопросов
	position := h.Tracker.GetCurrentPosition(currentTime)

	// Запускаем горутины для каждого вопроса (Fan-out)
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(questionNum int) {
			defer wg.Done()

			// Применяем rate limiter
			h.RateLimiter.Wait()

			// Применяем semaphore
			h.Semaphore.Acquire()
			defer h.Semaphore.Release()

			// Получаем воркера из load balancer
			worker := h.LoadBalancer.GetNextWorker()
			defer h.LoadBalancer.ReleaseWorker(worker)

			// Обрабатываем вопрос
			result := h.processQuestion(questionNum, currentTime, position, worker.ID)
			results <- result
		}(i)
	}

	// Ждём завершения всех горутин
	go func() {
		wg.Wait()
		close(results)
	}()

	// Fan-in: собираем результаты из канала
	allResults := make([]models.QuestionResult, 0, 10)
	for result := range results {
		allResults = append(allResults, result)
	}

	return allResults
}

// processQuestion обрабатывает конкретный вопрос
func (h *QuestionHandler) processQuestion(
	questionNum int,
	currentTime time.Time,
	position *models.CurrentPosition,
	workerID int,
) models.QuestionResult {
	h.Tracker.IncrementQuestionCounter(questionNum)

	result := models.QuestionResult{
		QuestionNumber: questionNum,
		ProcessedAt:    time.Now(),
	}

	switch questionNum {
	case 1:
		result.QuestionText = "Какое сейчас локальное время у пассажира?"
		result.Answer = h.Question1_LocalTime(currentTime, position)
	case 2:
		result.QuestionText = "На какой станции пассажир сейчас находится?"
		result.Answer = h.Question2_CurrentStation(position)
	case 3:
		result.QuestionText = "Поезд стоит или в пути?"
		result.Answer = h.Question3_TrainStatus(currentTime, position)
	case 4:
		result.QuestionText = "Какой день путешествия?"
		result.Answer = h.Question4_JourneyDay(currentTime)
	case 5:
		result.QuestionText = "Какое расстояние от Москвы?"
		result.Answer = h.Question5_Distance(position)
	case 6:
		result.QuestionText = "Когда пассажир прибудет на следующую станцию?"
		result.Answer = h.Question6_NextArrival(position)
	case 7:
		result.QuestionText = "Какая разница во времени между Москвой и текущим городом?"
		result.Answer = h.Question7_TimeDifference(currentTime, position)
	case 8:
		result.QuestionText = "Если я пишу сейчас, когда она получит?"
		result.Answer = h.Question8_MessageToHer(currentTime, position)
	case 9:
		result.QuestionText = "Если она пишет сейчас, когда я получу?"
		result.Answer = h.Question9_MessageFromHer(currentTime, position)
	case 10:
		result.QuestionText = "Какие основные станции впереди и когда прибытие?"
		result.Answer = h.Question10_UpcomingStations(position)
	}

	return result
}

// Question1_LocalTime - Какое сейчас локальное время у пассажира?
func (h *QuestionHandler) Question1_LocalTime(currentTime time.Time, pos *models.CurrentPosition) map[string]interface{} {
	if pos == nil {
		return map[string]interface{}{"error": "Position not found"}
	}

	localTime, _ := utils.ConvertToTimezone(currentTime, pos.Timezone)

	return map[string]interface{}{
		"local_time": localTime.Format("15:04 02.01.2006"),
		"timezone":   pos.Timezone,
	}
}

// Question2_CurrentStation - На какой станции пассажир сейчас находится?
func (h *QuestionHandler) Question2_CurrentStation(pos *models.CurrentPosition) map[string]interface{} {
	if pos == nil {
		return map[string]interface{}{"error": "Position not found"}
	}

	if pos.IsAtStation && pos.CurrentStation != nil {
		return map[string]interface{}{
			"station":          pos.CurrentStation.Name,
			"distance_from_moscow": pos.CurrentStation.DistanceFromStart,
			"at_station":       true,
		}
	}

	return map[string]interface{}{
		"between_stations": true,
		"previous":         pos.PreviousStation.Name,
		"next":             pos.NextStation.Name,
		"distance_from_moscow": int(pos.DistanceFromStart),
	}
}

// Question3_TrainStatus - Поезд стоит или в пути?
func (h *QuestionHandler) Question3_TrainStatus(currentTime time.Time, pos *models.CurrentPosition) map[string]interface{} {
	status := h.Tracker.GetTrainStatus(currentTime, pos)

	if !status.IsMoving {
		return map[string]interface{}{
			"status":            "СТОИТ",
			"station":           pos.CurrentStation.Name,
			"stand_duration":    utils.FormatDuration(pos.CurrentStation.StandDuration),
			"remaining_stand":   utils.FormatDuration(status.RemainingStand),
		}
	}

	return map[string]interface{}{
		"status":         "В ПУТИ",
		"from":           pos.PreviousStation.Name,
		"to":             pos.NextStation.Name,
		"time_to_next":   utils.FormatDuration(status.TimeToNext),
	}
}

// Question4_JourneyDay - Какой день путешествия?
func (h *QuestionHandler) Question4_JourneyDay(currentTime time.Time) map[string]interface{} {
	info := h.Tracker.GetJourneyInfo(currentTime)

	return map[string]interface{}{
		"day_number":       info.DayNumber,
		"start_date":       info.StartDate.Format("15:04 02.01.2006"),
		"time_in_trip":     utils.FormatDuration(info.TotalTimeInTrip),
	}
}

// Question5_Distance - Какое расстояние от Москвы?
func (h *QuestionHandler) Question5_Distance(pos *models.CurrentPosition) map[string]interface{} {
	if pos == nil {
		return map[string]interface{}{"error": "Position not found"}
	}

	location := "между станциями"
	if pos.IsAtStation && pos.CurrentStation != nil {
		location = pos.CurrentStation.Name
	}

	return map[string]interface{}{
		"distance_km": int(pos.DistanceFromStart),
		"location":    location,
	}
}

// Question6_NextArrival - Когда пассажир прибудет на следующую станцию?
// func (h *QuestionHandler) Question6_NextArrival(pos *models.CurrentPosition) map[string]interface{} {
// 	if pos == nil || pos.NextStation == nil {
// 		return map[string]interface{}{"error": "Next station not found"}
// 	}

// 	timeToNext := pos.NextStation.ArrivalTime.Sub(time.Now())

// 	return map[string]interface{}{
// 		"next_station":    pos.NextStation.Name,
// 		"arrival_time":    pos.NextStation.ArrivalTime.Format("15:04 02.01.2006"),
// 		"time_remaining":  utils.FormatDuration(timeToNext),
// 	}
// }
func (h *QuestionHandler) Question6_NextArrival(pos *models.CurrentPosition) map[string]interface{} {
    if pos == nil || pos.NextStation == nil {
        return map[string]interface{}{"error": "Next station not found"}
    }

    timeToNext := pos.NextStation.ArrivalTime.Sub(time.Now())
    
    // Handle negative time (train is late or algorithm issue)
    if timeToNext < 0 {
        timeToNext = 0
    }

    return map[string]interface{}{
        "next_station":    pos.NextStation.Name,
        "arrival_time":    pos.NextStation.ArrivalTime.Format("15:04 02.01.2006"),
        "time_remaining":  utils.FormatDuration(timeToNext),
    }
}

// Question7_TimeDifference - Какая разница во времени между Москвой и текущим городом?
func (h *QuestionHandler) Question7_TimeDifference(currentTime time.Time, pos *models.CurrentPosition) map[string]interface{} {
	if pos == nil {
		return map[string]interface{}{"error": "Position not found"}
	}

	moscowTime, _ := utils.ConvertToTimezone(currentTime, "Europe/Moscow")
	localTime, _ := utils.ConvertToTimezone(currentTime, pos.Timezone)

	diff, _ := utils.GetTimezoneDifference("Europe/Moscow", pos.Timezone)

	direction := "впереди Москвы"
	if diff < 0 {
		direction = "отстаёт от Москвы"
		diff = -diff
	}

	return map[string]interface{}{
		"moscow_time":      moscowTime.Format("15:04"),
		"local_time":       localTime.Format("15:04"),
		"difference":       utils.FormatDuration(diff),
		"direction":        direction,
	}
}

// Question8_MessageToHer - Если я пишу сейчас, когда она получит?
func (h *QuestionHandler) Question8_MessageToHer(currentTime time.Time, pos *models.CurrentPosition) map[string]interface{} {
	if pos == nil {
		return map[string]interface{}{"error": "Position not found"}
	}

	moscowTime, _ := utils.ConvertToTimezone(currentTime, "Europe/Moscow")
	herTime, _ := utils.ConvertToTimezone(currentTime, pos.Timezone)

	return map[string]interface{}{
		"send_time_moscow":   moscowTime.Format("15:04"),
		"receive_time_local": herTime.Format("15:04"),
		"instant_delivery":   true,
		"note":               "Сообщение доставляется мгновенно!",
	}
}

// Question9_MessageFromHer - Если она пишет сейчас, когда я получу?
func (h *QuestionHandler) Question9_MessageFromHer(currentTime time.Time, pos *models.CurrentPosition) map[string]interface{} {
	if pos == nil {
		return map[string]interface{}{"error": "Position not found"}
	}

	herTime, _ := utils.ConvertToTimezone(currentTime, pos.Timezone)
	moscowTime, _ := utils.ConvertToTimezone(currentTime, "Europe/Moscow")

	return map[string]interface{}{
		"send_time_local":     herTime.Format("15:04"),
		"receive_time_moscow": moscowTime.Format("15:04"),
		"instant_delivery":    true,
		"note":                "Сообщение доставляется мгновенно!",
	}
}

// Question10_UpcomingStations - Какие основные станции впереди и когда прибытие?
func (h *QuestionHandler) Question10_UpcomingStations(pos *models.CurrentPosition) map[string]interface{} {
	if pos == nil {
		return map[string]interface{}{"error": "Position not found"}
	}

	upcoming := []map[string]interface{}{}
	
	// Находим текущую позицию в массиве станций
	currentIndex := 0
	if pos.IsAtStation && pos.CurrentStation != nil {
		currentIndex = tracker.FindStationIndex(h.Tracker.Stations, pos.CurrentStation.ID)
	} else if pos.NextStation != nil {
		currentIndex = tracker.FindStationIndex(h.Tracker.Stations, pos.NextStation.ID)
	}

	// Берём только основные станции впереди
	count := 0
	for i := currentIndex; i < len(h.Tracker.Stations) && count < 10; i++ {
		station := h.Tracker.Stations[i]
		if station.IsMajor {
			upcoming = append(upcoming, map[string]interface{}{
				"name":          station.Name,
				"arrival_time":  station.ArrivalTime.Format("15:04 02.01.2006"),
				"stand_duration": utils.FormatDuration(station.StandDuration),
				"distance":      station.DistanceFromStart,
			})
			count++
		}
	}

	return map[string]interface{}{
		"upcoming_stations": upcoming,
		"count":             len(upcoming),
	}
}

// PrintResults красиво выводит результаты
func PrintResults(results []models.QuestionResult) {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("🔍 ОТВЕТЫ НА ВОПРОСЫ:")
	fmt.Println(strings.Repeat("=", 80))

	for _, result := range results {
		fmt.Printf("\n%d️⃣  %s\n", result.QuestionNumber, result.QuestionText)
		
		if answerMap, ok := result.Answer.(map[string]interface{}); ok {
			for key, value := range answerMap {
				fmt.Printf("   %s: %v\n", key, value)
			}
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
}

func (h *QuestionHandler) ProcessAllQuestionsWithRetry(currentTime time.Time) []models.QuestionResult {
    if h.Config != nil && h.Config.DebugMode {
        fmt.Println("🔄 Используется улучшенная обработка с повторными попытками...")
    }
    return h.enhancedProcessAllQuestions(currentTime)
}