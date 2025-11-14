package tracker

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"reyna-train-tracker/internal/cache"
	"reyna-train-tracker/internal/models"
	"reyna-train-tracker/internal/utils"
)

// TrainTracker основная структура для отслеживания поезда
type TrainTracker struct {
	Stations         []models.StationInfo
	StationsByName   map[string]*models.StationInfo // Hash table для быстрого доступа
	StationsByID     map[int]*models.StationInfo    // Hash table для быстрого доступа
	RouteData        models.RouteData
	Cache            *cache.InMemoryCache[interface{}] // In-memory cache с generic типом
	RequestCounter   atomic.Uint64                     // Atomic counter для статистики запросов
	QuestionCounters [11]atomic.Uint64                 // Счётчики для каждого из 10 вопросов (индекс 0 не используется)
}

// NewTrainTracker создаёт новый трекер
func NewTrainTracker(jsonPath string) (*TrainTracker, error) {
	tracker := &TrainTracker{
		Cache: cache.NewInMemoryCache[interface{}](),
	}

	err := tracker.LoadSchedule(jsonPath)
	if err != nil {
		return nil, err
	}

	// Строим hash tables для быстрого доступа
	tracker.StationsByName, tracker.StationsByID = BuildStationHashMap(tracker.Stations)

	return tracker, nil
}

// LoadSchedule загружает расписание из JSON файла
// func (t *TrainTracker) LoadSchedule(jsonPath string) error {
// 	data, err := os.ReadFile(jsonPath)
// 	if err != nil {
// 		return fmt.Errorf("failed to read JSON file: %w", err)
// 	}

// 	// Парсим JSON
// 	var rawData map[string]models.Station
// 	err = json.Unmarshal(data, &rawData)
// 	if err != nil {
// 		return fmt.Errorf("failed to unmarshal JSON: %w", err)
// 	}

// 	// Преобразуем в StationInfo с полной информацией
// 	// Начальная дата - 6 октября 2025, время отправления из Москвы - 22:10
// 	moscowTZ, _ := time.LoadLocation("Europe/Moscow")
// 	startDate := time.Date(2025, 10, 6, 22, 10, 0, 0, moscowTZ)

// 	t.RouteData = models.RouteData{
// 		Name:      "Москва - Хабаровск",
// 		StartTime: startDate,
// 	}

// 	// Сортируем станции по ID
// 	sortedKeys := make([]string, 0, len(rawData))
// 	for key := range rawData {
// 		sortedKeys = append(sortedKeys, key)
// 	}
// 	sort.Strings(sortedKeys)

// 	currentDate := startDate
// 	currentDeparture := startDate

// 	// Обрабатываем станции по порядку
// 	for _, key := range sortedKeys {
// 		station := rawData[key]
		
// 		// Парсим номер станции
// 		stationID, err := ParseCityNumber(key)
// 		if err != nil {
// 			continue
// 		}

// 		stationInfo := models.StationInfo{
// 			ID:   stationID,
// 			Name: station.Name,
// 		}

// 		// Получаем часовой пояс
// 		stationInfo.Timezone = utils.GetTimezone(station.Name)

// 		// Парсим длительность стоянки
// 		standDuration, _ := utils.ParseStandDuration(station.Stand)
// 		stationInfo.StandDuration = standDuration

// 		// Парсим время прибытия
// 		arrivalTime, err := utils.ParseTime(station.TimeArrive, currentDate)
// 		if err != nil {
// 			fmt.Printf("❌ Ошибка парсинга времени прибытия для %s: %v\n", station.Name, err)
// 			continue
// 		}

// 		// Если время прибытия РАНЬШЕ времени отправления предыдущей станции, добавляем день
// 		if arrivalTime.Before(currentDeparture) {
// 			currentDate = currentDate.Add(24 * time.Hour)
// 			arrivalTime, _ = utils.ParseTime(station.TimeArrive, currentDate)
// 		}
// 		stationInfo.ArrivalTime = arrivalTime

// 		// Парсим время отправления
// 		departureTime, err := utils.ParseTime(station.TimeDepart, currentDate)
// 		if err != nil {
// 			fmt.Printf("❌ Ошибка парсинга времени отправления для %s: %v\n", station.Name, err)
// 			continue
// 		}

// 		// Если время отправления РАНЬШЕ времени прибытия, добавляем день
// 		if departureTime.Before(arrivalTime) {
// 			currentDate = currentDate.Add(24 * time.Hour)
// 			departureTime, _ = utils.ParseTime(station.TimeDepart, currentDate)
// 		}
// 		stationInfo.DepartureTime = departureTime
// 		currentDeparture = departureTime

// 		// Получаем расстояние
// 		stationInfo.DistanceFromStart = utils.GetDistance(station.Name)

// 		// Определяем основные станции
// 		stationInfo.IsMajor = standDuration >= 20*time.Minute || isMajorCity(station.Name)

// 		t.Stations = append(t.Stations, stationInfo)
		
// 		// Отладочный вывод для первых нескольких станций
// 		if stationID <= 5 {
// 			fmt.Printf("🚉 Station %d: %s\n", stationID, station.Name)
// 			fmt.Printf("   Arrival: %s | Departure: %s\n", 
// 				stationInfo.ArrivalTime.Format("15:04 02.01"), 
// 				stationInfo.DepartureTime.Format("15:04 02.01"))
// 		}
// 	}

// 	// Устанавливаем общее расстояние
// 	if len(t.Stations) > 0 {
// 		t.RouteData.TotalDistance = t.Stations[len(t.Stations)-1].DistanceFromStart
// 	}

// 	return nil
// }

// LoadSchedule загружает расписание из JSON файла
func (t *TrainTracker) LoadSchedule(jsonPath string) error {
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return fmt.Errorf("failed to read JSON file: %w", err)
	}

	// Парсим JSON
	var rawData map[string]models.Station
	err = json.Unmarshal(data, &rawData)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Преобразуем в StationInfo с полной информацией
	moscowTZ, _ := time.LoadLocation("Europe/Moscow")
	currentDate := time.Date(2025, 10, 6, 0, 0, 0, 0, moscowTZ) // Начинаем с 6 октября

	t.RouteData = models.RouteData{
		Name:      "Москва - Хабаровск", 
		StartTime: time.Date(2025, 10, 6, 22, 10, 0, 0, moscowTZ),
	}

	// Сортируем станции
	sortedKeys := make([]string, 0, len(rawData))
	for key := range rawData {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)

	fmt.Printf("🔍 КОРРЕКТИРОВАННАЯ ЗАГРУЗКА С 6 ПО 13 ОКТЯБРЯ:\n")

	// Обрабатываем станции по порядку
	for _, key := range sortedKeys {
		station := rawData[key]
		
		stationID, err := ParseCityNumber(key)
		if err != nil {
			continue
		}

		stationInfo := models.StationInfo{
			ID:   stationID,
			Name: station.Name,
		}

		// Получаем часовой пояс
		stationInfo.Timezone = utils.GetTimezone(station.Name)

		// Парсим длительность стоянки
		standDuration, _ := utils.ParseStandDuration(station.Stand)
		stationInfo.StandDuration = standDuration

		// Парсим время прибытия (в московском времени)
		arrivalTime, err := utils.ParseTime(station.TimeArrive, currentDate)
		if err != nil {
			continue
		}

		// Парсим время отправления  
		departureTime, err := utils.ParseTime(station.TimeDepart, currentDate)
		if err != nil {
			continue
		}

		if departureTime.Before(arrivalTime) {
			fmt.Printf("⚠️  ИСПРАВЛЕНО: %s - отправление раньше прибытия\n", station.Name)
			departureTime = arrivalTime.Add(5 * time.Minute) // Минимальная стоянка 5 минут
		}

		// КОРРЕКТНАЯ ЛОГИКА: добавляем день только если время меньше предыдущего
		// (это означает переход через полночь)
		if stationID > 1 {
			prevStation := t.Stations[stationID-2] // предыдущая станция
			
			// Если прибытие раньше отправления предыдущей станции - следующий день
			if arrivalTime.Before(prevStation.DepartureTime) {
				currentDate = currentDate.Add(24 * time.Hour)
				arrivalTime, _ = utils.ParseTime(station.TimeArrive, currentDate)
				departureTime, _ = utils.ParseTime(station.TimeDepart, currentDate)
			}
			
			// Если отправление раньше прибытия - следующий день
			if departureTime.Before(arrivalTime) {
				currentDate = currentDate.Add(24 * time.Hour)
				departureTime, _ = utils.ParseTime(station.TimeDepart, currentDate)
			}
		}

		stationInfo.ArrivalTime = arrivalTime
		stationInfo.DepartureTime = departureTime

		// Получаем расстояние
		stationInfo.DistanceFromStart = utils.GetDistance(station.Name)

		// Определяем основные станции
		stationInfo.IsMajor = standDuration >= 20*time.Minute || isMajorCity(station.Name)

		t.Stations = append(t.Stations, stationInfo)
		
		// Выводим ВСЕ станции для проверки
		fmt.Printf("🚉 %2d: %-30s | %s - %s | %s\n", 
			stationID, station.Name,
			arrivalTime.Format("15:04 02.01"),
			departureTime.Format("15:04 02.01"),
			stationInfo.Timezone)
	}

	// Проверяем дату прибытия в Хабаровск
	if len(t.Stations) > 0 {
		lastStation := t.Stations[len(t.Stations)-1]
		fmt.Printf("\n📅 ПРИБЫТИЕ В ХАБАРОВСК: %s\n", 
			lastStation.ArrivalTime.Format("15:04 02.01.2006"))
		
		t.RouteData.TotalDistance = lastStation.DistanceFromStart
	}

	return nil
}

// DebugAllStations отладочная функция для вывода всех станций
func (t *TrainTracker) DebugAllStations() {
	fmt.Printf("\n🔍 DEBUG ALL STATIONS TIMELINE:\n")
	for i, station := range t.Stations {
		if i < 10 || i > len(t.Stations)-10 { // Show first and last 10 stations
			fmt.Printf("Station %2d: %-30s | Arr: %s | Dep: %s | Dist: %dkm\n",
				station.ID, station.Name,
				station.ArrivalTime.Format("15:04 02.01"),
				station.DepartureTime.Format("15:04 02.01"),
				station.DistanceFromStart)
		} else if i == 10 {
			fmt.Printf("... (middle stations omitted)\n")
		}
	}
}

// isMajorCity определяет, является ли город крупным
func isMajorCity(name string) bool {
	majorCities := []string{
		"Москва",
		"Владимир Пасс",
		"Нижний Новгород Московский (Московский вокзал)",
		"Киров Пасс",
		"Пермь 2",
		"Екатеринбург-Пассажирс",
		"Тюмень",
		"Омск-Пассажирский",
		"Новосибирск-Главный",
		"Красноярск Пасс",
		"Иркутск Пассажирский",
		"Улан-Удэ Пасс",
		"Чита 2",
		"Сковородино",
		"Биробиджан 1",
		"Хабаровск 1",
	}

	for _, major := range majorCities {
		if name == major {
			return true
		}
	}
	return false
}

// GetStationByName получает станцию по названию (O(1) благодаря hash table)
func (t *TrainTracker) GetStationByName(name string) (*models.StationInfo, bool) {
	station, ok := t.StationsByName[name]
	return station, ok
}

// GetStationByID получает станцию по ID (O(1) благодаря hash table)
func (t *TrainTracker) GetStationByID(id int) (*models.StationInfo, bool) {
	station, ok := t.StationsByID[id]
	return station, ok
}

// GetCurrentPosition получает текущую позицию пассажира
// Использует алгоритм двух указателей и кэш
func (t *TrainTracker) GetCurrentPosition(currentTime time.Time) *models.CurrentPosition {
	// Увеличиваем счётчик запросов (atomic operation)
	t.RequestCounter.Add(1)

	// Проверяем кэш
	cacheKey := fmt.Sprintf("position_%d", currentTime.Unix())
	if cached, ok := t.Cache.Get(cacheKey); ok {
		if pos, ok := cached.(*models.CurrentPosition); ok {
			return pos
		}
	}

	// Используем алгоритм двух указателей
	pos := FindCurrentPositionTwoPointers(t.Stations, currentTime)

	if pos != nil {
		// Конвертируем время в локальный часовой пояс
		localTime, _ := utils.ConvertToTimezone(currentTime, pos.Timezone)
		pos.LocalTime = localTime

		// Кэшируем результат на 1 минуту
		t.Cache.Set(cacheKey, pos, 1*time.Minute)
	}

	return pos
}

// GetTrainStatus получает статус поезда (стоит или едет)
func (t *TrainTracker) GetTrainStatus(currentTime time.Time, pos *models.CurrentPosition) models.TrainStatus {
	status := models.TrainStatus{}

	if pos == nil {
		return status
	}

	if pos.IsAtStation && pos.CurrentStation != nil {
		// Поезд стоит на станции
		status.IsMoving = false
		timeUntilDeparture := pos.CurrentStation.DepartureTime.Sub(currentTime)
		if timeUntilDeparture > 0 {
			status.RemainingStand = timeUntilDeparture
		}
	} else if pos.NextStation != nil {
		// Поезд в движении
		status.IsMoving = true
		status.TimeToNext = pos.NextStation.ArrivalTime.Sub(currentTime)
	}

	return status
}

// GetJourneyInfo получает информацию о путешествии
func (t *TrainTracker) GetJourneyInfo(currentTime time.Time) models.JourneyInfo {
	info := models.JourneyInfo{
		StartDate: t.RouteData.StartTime,
	}

	// Рассчитываем общее время в пути
	info.TotalTimeInTrip = currentTime.Sub(t.RouteData.StartTime)

	// Рассчитываем день путешествия
	days := int(info.TotalTimeInTrip.Hours() / 24)
	info.DayNumber = days + 1

	return info
}

// IncrementQuestionCounter увеличивает счётчик для конкретного вопроса
func (t *TrainTracker) IncrementQuestionCounter(questionNumber int) {
	if questionNumber >= 1 && questionNumber <= 10 {
		t.QuestionCounters[questionNumber].Add(1)
	}
}

// GetStatistics возвращает статистику использования
func (t *TrainTracker) GetStatistics() map[string]interface{} {
	stats := make(map[string]interface{})
	stats["total_requests"] = t.RequestCounter.Load()

	questionStats := make(map[string]uint64)
	for i := 1; i <= 10; i++ {
		questionStats[fmt.Sprintf("question_%d", i)] = t.QuestionCounters[i].Load()
	}
	stats["question_counters"] = questionStats
	stats["cache_size"] = t.Cache.Size()

	return stats
}

// ParseCityNumber парсит номер города из строки типа "city_38" или "city_0038"
func ParseCityNumber(cityKey string) (int, error) {
	var num int
	
	// Пробуем разные форматы
	_, err := fmt.Sscanf(cityKey, "city_%d", &num)
	if err != nil {
		// Пробуем zero-padded формат
		_, err = fmt.Sscanf(cityKey, "city_%04d", &num)
		if err != nil {
			// Альтернативный парсинг - берём часть после "city_"
			numStr := strings.TrimPrefix(cityKey, "city_")
			num, err = strconv.Atoi(numStr)
		}
	}
	return num, err
}
