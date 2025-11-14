package tracker

import (
	"fmt"
	"time"

	"reyna-train-tracker/internal/models"
)

// ImprovedTwoPointersSearch - улучшенный алгоритм поиска с бинарным поиском
func ImprovedTwoPointersSearch(stations []models.StationInfo, currentTime time.Time) *models.CurrentPosition {
	if len(stations) == 0 {
		return nil
	}

	// Проверяем граничные случаи
	if currentTime.Before(stations[0].DepartureTime) {
		return createStationPosition(stations, 0, true)
	}

	if currentTime.After(stations[len(stations)-1].ArrivalTime) {
		return createStationPosition(stations, len(stations)-1, true)
	}

	// Бинарный поиск для нахождения ближайшей станции
	left, right := 0, len(stations)-1
	
	for left <= right {
		mid := left + (right-left)/2
		station := stations[mid]
		
		// Проверяем, находимся ли на станции
		if !currentTime.Before(station.ArrivalTime) && !currentTime.After(station.DepartureTime) {
			return createStationPosition(stations, mid, true)
		}
		
		// Проверяем, находимся ли между текущей и следующей станцией
		if mid < len(stations)-1 {
			nextStation := stations[mid+1]
			if currentTime.After(station.DepartureTime) && currentTime.Before(nextStation.ArrivalTime) {
				return createBetweenPosition(station, nextStation, currentTime)
			}
		}
		
		// Проверяем, находимся ли между предыдущей и текущей станцией
		if mid > 0 {
			prevStation := stations[mid-1]
			if currentTime.After(prevStation.DepartureTime) && currentTime.Before(station.ArrivalTime) {
				return createBetweenPosition(prevStation, station, currentTime)
			}
		}
		
		// Определяем направление поиска
		if currentTime.Before(station.ArrivalTime) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	
	// Fallback: линейный поиск для edge cases
	return fallbackLinearSearch(stations, currentTime)
}

func createStationPosition(stations []models.StationInfo, index int, atStation bool) *models.CurrentPosition {
	pos := &models.CurrentPosition{
		IsAtStation:       atStation,
		CurrentStation:    &stations[index],
		DistanceFromStart: float64(stations[index].DistanceFromStart),
		Timezone:          stations[index].Timezone,
	}
	
	if index > 0 {
		pos.PreviousStation = &stations[index-1]
	}
	if index < len(stations)-1 {
		pos.NextStation = &stations[index+1]
	}
	
	return pos
}

func createBetweenPosition(prev, next models.StationInfo, currentTime time.Time) *models.CurrentPosition {
	totalTime := next.ArrivalTime.Sub(prev.DepartureTime).Seconds()
	elapsed := currentTime.Sub(prev.DepartureTime).Seconds()
	
	if totalTime <= 0 {
		totalTime = 1 // Avoid division by zero
	}
	
	progress := clamp(elapsed/totalTime, 0, 1)
	currentDist := float64(prev.DistanceFromStart) + 
		(float64(next.DistanceFromStart-prev.DistanceFromStart)) * progress
	
	return &models.CurrentPosition{
		IsAtStation:       false,
		PreviousStation:   &prev,
		NextStation:       &next,
		DistanceFromStart: currentDist,
		Timezone:          prev.Timezone,
	}
}

func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func fallbackLinearSearch(stations []models.StationInfo, currentTime time.Time) *models.CurrentPosition {
	// Линейный поиск как fallback
	for i := 0; i < len(stations); i++ {
		station := stations[i]
		
		// Проверяем, находимся ли на станции
		if !currentTime.Before(station.ArrivalTime) && !currentTime.After(station.DepartureTime) {
			return createStationPosition(stations, i, true)
		}
		
		// Проверяем, находимся ли между станциями
		if i < len(stations)-1 {
			nextStation := stations[i+1]
			if currentTime.After(station.DepartureTime) && currentTime.Before(nextStation.ArrivalTime) {
				return createBetweenPosition(station, nextStation, currentTime)
			}
		}
	}
	
	return nil
}

// FindCurrentPositionTwoPointers - оригинальный алгоритм (оставляем для обратной совместимости)
func FindCurrentPositionTwoPointers(stations []models.StationInfo, currentTime time.Time) *models.CurrentPosition {
	// Используем улучшенный алгоритм по умолчанию
	return ImprovedTwoPointersSearch(stations, currentTime)
}

func DebugFindCurrentPosition(stations []models.StationInfo, currentTime time.Time) {
	fmt.Printf("\n🔍 DEBUG POSITION CALCULATION:\n")
	fmt.Printf("Current Time: %s\n", currentTime.Format("15:04 02.01.2006"))
	
	// Находим приблизительную позицию для отладки
	position := ImprovedTwoPointersSearch(stations, currentTime)
	if position != nil {
		if position.IsAtStation && position.CurrentStation != nil {
			fmt.Printf("📍 На станции: %s\n", position.CurrentStation.Name)
		} else if position.PreviousStation != nil && position.NextStation != nil {
			fmt.Printf("📍 Между станциями: %s -> %s\n", 
				position.PreviousStation.Name, position.NextStation.Name)
		}
	}
	
	// Показываем станции вокруг текущей позиции
	startIdx := 0
	if position != nil && position.PreviousStation != nil {
		startIdx = max(0, position.PreviousStation.ID-2)
	}
	
	endIdx := min(len(stations), startIdx+8)
	
	for i := startIdx; i < endIdx; i++ {
		if i < len(stations) {
			station := stations[i]
			fmt.Printf("Station %d: %s\n", station.ID, station.Name)
			fmt.Printf("  Arrival: %s | Departure: %s\n", 
				station.ArrivalTime.Format("15:04 02.01"), 
				station.DepartureTime.Format("15:04 02.01"))
			fmt.Printf("  Before arrival? %v | After departure? %v\n",
				currentTime.Before(station.ArrivalTime),
				currentTime.After(station.DepartureTime))
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Алгоритм 2: ХЭШ-ТАБЛИЦЫ (Hash Tables)
// Используется для быстрого O(1) доступа к данным станций

// BuildStationHashMap создаёт хэш-таблицу для быстрого доступа к станциям
// по названию и по ID
func BuildStationHashMap(stations []models.StationInfo) (map[string]*models.StationInfo, map[int]*models.StationInfo) {
	nameMap := make(map[string]*models.StationInfo)
	idMap := make(map[int]*models.StationInfo)

	for i := range stations {
		station := &stations[i]
		nameMap[station.Name] = station
		idMap[station.ID] = station
	}

	return nameMap, idMap
}

// Алгоритм 3: СКОЛЬЗЯЩЕЕ ОКНО (Sliding Window)
// Используется для предсказания времени прибытия и расчёта средней скорости

// CalculateAverageSpeedSlidingWindow использует скользящее окно
// для расчёта средней скорости на последних N отрезках
func CalculateAverageSpeedSlidingWindow(stations []models.StationInfo, currentIndex int, windowSize int) float64 {
	if currentIndex < 1 || len(stations) < 2 {
		// Средняя скорость поезда ~90 км/ч
		return 90.0
	}

	// Определяем размер окна
	start := currentIndex - windowSize
	if start < 0 {
		start = 0
	}

	totalDistance := 0.0
	totalTime := 0.0

	// Скользящее окно: берём последние windowSize отрезков
	for i := start; i < currentIndex && i < len(stations)-1; i++ {
		distance := float64(stations[i+1].DistanceFromStart - stations[i].DistanceFromStart)
		duration := stations[i+1].ArrivalTime.Sub(stations[i].DepartureTime).Hours()

		totalDistance += distance
		totalTime += duration
	}

	if totalTime > 0 {
		return totalDistance / totalTime
	}

	return 90.0 // Средняя скорость по умолчанию
}

// PredictArrivalTime предсказывает время прибытия на основе скользящего окна
func PredictArrivalTime(
	fromStation *models.StationInfo,
	toStation *models.StationInfo,
	currentTime time.Time,
	averageSpeed float64,
) time.Time {
	// Расстояние до следующей станции
	distance := float64(toStation.DistanceFromStart - fromStation.DistanceFromStart)

	// Время в пути = расстояние / скорость
	hoursNeeded := distance / averageSpeed

	// Предсказанное время прибытия
	return currentTime.Add(time.Duration(hoursNeeded * float64(time.Hour)))
}

// GetMajorStations возвращает основные станции (для вопроса 10)
func GetMajorStations(stations []models.StationInfo) []models.StationInfo {
	majorStations := []models.StationInfo{}

	for _, station := range stations {
		if station.IsMajor {
			majorStations = append(majorStations, station)
		}
	}

	return majorStations
}

// FindStationIndex находит индекс станции в массиве
func FindStationIndex(stations []models.StationInfo, stationID int) int {
	for i, station := range stations {
		if station.ID == stationID {
			return i
		}
	}
	return -1
}