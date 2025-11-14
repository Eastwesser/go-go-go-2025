package api

import (
	"fmt"
	"reyna-train-tracker/internal/models"
	"sync"
	"time"
)

// processQuestionWithRetry обрабатывает вопрос с повторными попытками при ошибках
func (h *QuestionHandler) processQuestionWithRetry(
    questionNum int,
    currentTime time.Time,
    position *models.CurrentPosition,
    workerID int,
    maxRetries int,
) models.QuestionResult {
    var result models.QuestionResult
    var lastErr error
    
    for attempt := 0; attempt < maxRetries; attempt++ {
        startTime := time.Now()
        result = h.processQuestion(questionNum, currentTime, position, workerID)
        processingTime := time.Since(startTime)
        
        // Записываем метрику
        if h.Metrics != nil {
            h.Metrics.RecordRequest(processingTime, true)
        }
        
        // Проверяем наличие ошибок в ответе
        if answerMap, ok := result.Answer.(map[string]interface{}); ok {
            if _, hasError := answerMap["error"]; !hasError {
                // Успешная обработка
                if h.Metrics != nil {
                    h.Metrics.RecordCacheHit()
                }
                return result
            }
            lastErr = fmt.Errorf("attempt %d: %v", attempt+1, answerMap["error"])
            
            // Записываем ошибку в метрики
            if h.Metrics != nil {
                h.Metrics.RecordRequest(processingTime, false)
            }
        }
        
        // Exponential backoff перед следующей попыткой
        if attempt < maxRetries-1 {
            backoffDuration := time.Duration(attempt+1) * 100 * time.Millisecond
            if h.Config != nil && h.Config.DebugMode {
                fmt.Printf("🔄 Повторная попытка %d/%d для вопроса %d через %v\n", 
                    attempt+1, maxRetries, questionNum, backoffDuration)
            }
            time.Sleep(backoffDuration)
        }
    }
    
    // Если все попытки неудачны, возвращаем ошибку
    if lastErr != nil {
        result.Answer = map[string]interface{}{
            "error": fmt.Sprintf("❌ Не удалось обработать вопрос после %d попыток: %v", maxRetries, lastErr),
            "question_number": questionNum,
            "max_retries": maxRetries,
        }
    }
    
    return result
}

// enhancedProcessAllQuestions улучшенная версия обработки всех вопросов с retry логикой
func (h *QuestionHandler) enhancedProcessAllQuestions(currentTime time.Time) []models.QuestionResult {
    // Получаем текущую позицию один раз для всех вопросов
    position := h.Tracker.GetCurrentPosition(currentTime)
    
    // Используем конфигурацию для определения количества повторов
    maxRetries := 3 // значение по умолчанию
    if h.Config != nil {
        maxRetries = h.Config.MaxRetries
    }

    results := make(chan models.QuestionResult, 10)
    var wg sync.WaitGroup

    // Fan-out: запускаем обработку всех вопросов параллельно с retry
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

            // Обрабатываем вопрос с повторными попытками
            result := h.processQuestionWithRetry(questionNum, currentTime, position, worker.ID, maxRetries)
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

// validateQuestionResult проверяет валидность результата вопроса
func (h *QuestionHandler) validateQuestionResult(result models.QuestionResult) bool {
    if answerMap, ok := result.Answer.(map[string]interface{}); ok {
        // Проверяем наличие ошибки
        if _, hasError := answerMap["error"]; hasError {
            return false
        }
        
        // Валидация в зависимости от типа вопроса
        switch result.QuestionNumber {
        case 1: // Локальное время
            return answerMap["local_time"] != nil && answerMap["timezone"] != nil
        case 2: // Текущая станция
            return answerMap["distance_from_moscow"] != nil
        case 3: // Статус поезда
            return answerMap["status"] != nil
        case 4: // День путешествия
            return answerMap["day_number"] != nil
        case 5: // Расстояние
            return answerMap["distance_km"] != nil
        case 6: // Следующая станция
            return answerMap["next_station"] != nil && answerMap["arrival_time"] != nil
        case 7: // Разница во времени
            return answerMap["difference"] != nil
        case 8, 9: // Сообщения
            return answerMap["send_time_moscow"] != nil || answerMap["send_time_local"] != nil
        case 10: // Основные станции
            return answerMap["upcoming_stations"] != nil
        }
    }
    return false
}

// getQuestionRetryStats возвращает статистику по повторным попыткам
func (h *QuestionHandler) getQuestionRetryStats(results []models.QuestionResult) map[string]interface{} {
    stats := map[string]interface{}{
        "total_questions": len(results),
        "successful": 0,
        "failed": 0,
        "retry_attempts": make(map[int]int),
    }
    
    for _, result := range results {
        if h.validateQuestionResult(result) {
            stats["successful"] = stats["successful"].(int) + 1
        } else {
            stats["failed"] = stats["failed"].(int) + 1
        }
        
        // Анализируем ответы на наличие информации о retry
        if answerMap, ok := result.Answer.(map[string]interface{}); ok {
            if attempts, exists := answerMap["max_retries"]; exists {
                questionNum := result.QuestionNumber
                stats["retry_attempts"].(map[int]int)[questionNum] = attempts.(int)
            }
        }
    }
    
    return stats
}