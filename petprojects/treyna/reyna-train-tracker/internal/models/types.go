package models

import (
	"time"
)

// Station представляет базовую информацию о станции из JSON
type Station struct {
	Name        string `json:"name"`
	TimeArrive  string `json:"timeArrive"`
	Stand       string `json:"stand"`
	TimeDepart  string `json:"timeDepart"`
}

// StationInfo расширенная информация о станции с расчётами
type StationInfo struct {
	ID                int           // ID станции (city_2 = 2)
	Name              string        // Название станции
	Timezone          string        // Часовой пояс станции
	ArrivalTime       time.Time     // Время прибытия (в московском времени)
	DepartureTime     time.Time     // Время отправления (в московском времени)
	StandDuration     time.Duration // Длительность стоянки
	DistanceFromStart int           // Расстояние от Москвы в км (приблизительное)
	IsMajor           bool          // Основная станция (для вопроса 10)
}

// RouteData содержит информацию о маршруте
type RouteData struct {
	Name          string
	TotalDistance int
	StartTime     time.Time
	Stations      []StationInfo
}

// CurrentPosition текущая позиция пассажира
type CurrentPosition struct {
	IsAtStation       bool          // На станции или между станциями
	CurrentStation    *StationInfo  // Текущая станция (если на станции)
	PreviousStation   *StationInfo  // Предыдущая станция
	NextStation       *StationInfo  // Следующая станция
	DistanceFromStart float64       // Расстояние от Москвы в км
	LocalTime         time.Time     // Локальное время пассажира
	Timezone          string        // Текущий часовой пояс
}

// TrainStatus статус поезда
type TrainStatus struct {
	IsMoving        bool
	RemainingStand  time.Duration // Осталось стоять (если стоит)
	TimeToNext      time.Duration // Время до следующей станции (если движется)
}

// JourneyInfo информация о путешествии
type JourneyInfo struct {
	DayNumber        int           // Какой день путешествия
	StartDate        time.Time     // Дата начала путешествия
	TotalTimeInTrip  time.Duration // Общее время в пути
}

// MessageDelivery информация о доставке сообщений
type MessageDelivery struct {
	SenderTime   time.Time
	ReceiverTime time.Time
	Instant      bool // Мгновенная доставка
}

// CacheEntry запись в кэше
type CacheEntry[T any] struct {
	Value     T
	Timestamp time.Time
	TTL       time.Duration
}

// QuestionResult результат ответа на вопрос
type QuestionResult struct {
	QuestionNumber int
	QuestionText   string
	Answer         interface{}
	ProcessedAt    time.Time
}

