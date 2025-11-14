package utils

import (
	"fmt"
	"strings"
	"time"
)

// TimezoneMap маппинг городов на часовые пояса
var TimezoneMap = map[string]string{
	// Europe/Moscow (UTC+3)
	"Москва":                                     "Europe/Moscow",
	"Владимир Пасс":                              "Europe/Moscow",
	"Ковров 1":                                   "Europe/Moscow",
	"Нижний Новгород Московский (Московский вокзал)": "Europe/Moscow",
	"Семенов":                                    "Europe/Moscow",
	"Киров Пасс":                                 "Europe/Moscow",
	"Зуевка":                                     "Europe/Moscow",

	// Asia/Yekaterinburg (UTC+5)
	"Глазов":              "Asia/Yekaterinburg",
	"Балезино":            "Asia/Yekaterinburg",
	"Пермь 2":             "Asia/Yekaterinburg",
	"Екатеринбург-Пассажирс": "Asia/Yekaterinburg",
	"Тюмень":              "Asia/Yekaterinburg",

	// Asia/Omsk (UTC+6)
	"Омск-Пассажирский":   "Asia/Omsk",
	"Татарская":           "Asia/Omsk",
	"Озеро-Карачинское":   "Asia/Omsk",
	"Барабинск":           "Asia/Omsk",

	// Asia/Novosibirsk (UTC+7)
	"Новосибирск-Главный": "Asia/Novosibirsk",
	"Юрга 1":              "Asia/Krasnoyarsk",
	"Яшкино":              "Asia/Krasnoyarsk",
	"Тайга":               "Asia/Krasnoyarsk",
	"Анжерская":           "Asia/Krasnoyarsk",
	"Яя":                  "Asia/Krasnoyarsk",
	"Мариинск":            "Asia/Krasnoyarsk",
	"Тяжин":               "Asia/Krasnoyarsk",
	"Боготол":             "Asia/Krasnoyarsk",
	"Ачинск 1":            "Asia/Krasnoyarsk",
	"Красноярск Пасс":     "Asia/Krasnoyarsk",
	"Уяр":                 "Asia/Krasnoyarsk",
	"Заозерная":           "Asia/Krasnoyarsk",
	"Канск-Енисейский":    "Asia/Krasnoyarsk",
	"Иланская":            "Asia/Krasnoyarsk",
	"Ингашская":           "Asia/Krasnoyarsk",
	"Решоты":              "Asia/Krasnoyarsk",
	"Юрты":                "Asia/Krasnoyarsk",
	"Тайшет":              "Asia/Krasnoyarsk",
	"Нижнеудинск":         "Asia/Krasnoyarsk",

	// Asia/Irkutsk (UTC+8)
	"Тулун":                "Asia/Irkutsk",
	"Зима":                 "Asia/Irkutsk",
	"Залари":               "Asia/Irkutsk",
	"Черемхово":            "Asia/Irkutsk",
	"Усолье-Сибирское":     "Asia/Irkutsk",
	"Ангарск":              "Asia/Irkutsk",
	"Иркутск-Сорт":         "Asia/Irkutsk",
	"Иркутск Пассажирский": "Asia/Irkutsk",
	"Слюдянка 1":           "Asia/Irkutsk",
	"Байкальск":            "Asia/Irkutsk",
	"Мысовая":              "Asia/Irkutsk",
	"Улан-Удэ Пасс":        "Asia/Irkutsk",
	"Заудинский":           "Asia/Irkutsk",
	"Новоильинский":        "Asia/Irkutsk",
	"Петровский Завод":     "Asia/Irkutsk",
	"Бада":                 "Asia/Irkutsk",
	"Хилок":                "Asia/Irkutsk",
	"Хушенга":              "Asia/Irkutsk",
	"Харагун":              "Asia/Irkutsk",
	"Могзон":               "Asia/Irkutsk",

	// Asia/Yakutsk (UTC+9)
	"Чита 2":                   "Asia/Yakutsk",
	"Карымская":                "Asia/Yakutsk",
	"Солнцевая":                "Asia/Yakutsk",
	"Шилка-Пасс.":              "Asia/Yakutsk",
	"Приисковая":               "Asia/Yakutsk",
	"Куэнга":                   "Asia/Yakutsk",
	"Чернышевск-Забайкальск":   "Asia/Yakutsk",
	"Жирекен":                  "Asia/Yakutsk",
	"Зилово":                   "Asia/Yakutsk",
	"Ксеньевская ":             "Asia/Yakutsk",
	"Могоча":                   "Asia/Yakutsk",
	"Амазар":                   "Asia/Yakutsk",
	"Ерофей Павлович":          "Asia/Yakutsk",
	"Уруша":                    "Asia/Yakutsk",
	"Сковородино":              "Asia/Yakutsk",
	"Талдан":                   "Asia/Yakutsk",
	"Магдагачи":                "Asia/Yakutsk",
	"Тыгда":                    "Asia/Yakutsk",
	"Шимановская":              "Asia/Yakutsk",
	"Ледяная":                  "Asia/Yakutsk",
	"Свободный":                "Asia/Yakutsk",
	"Серышево":                 "Asia/Yakutsk",
	"Белогорск":                "Asia/Yakutsk",
	"Поздеевка":                "Asia/Yakutsk",
	"Екатеринославка":          "Asia/Yakutsk",
	"Завитая":                  "Asia/Yakutsk",
	"Бурея":                    "Asia/Yakutsk",
	"Архара":                   "Asia/Yakutsk",

	// Asia/Vladivostok (UTC+10)
	"Облучье":        "Asia/Vladivostok",
	"Известковая":    "Asia/Vladivostok",
	"Биробиджан 1":   "Asia/Vladivostok",
	"Хабаровск 1":    "Asia/Vladivostok",
}

// GetTimezone получает часовой пояс для города
func GetTimezone(cityName string) string {
	if tz, ok := TimezoneMap[cityName]; ok {
		return tz
	}
	// По умолчанию Москва
	return "Europe/Moscow"
}

// ParseStandDuration парсит длительность стоянки из строки типа "20мин", "1ч", "2мин"
func ParseStandDuration(stand string) (time.Duration, error) {
	stand = strings.TrimSpace(stand)

	// Удаляем пробелы
	stand = strings.ReplaceAll(stand, " ", "")

	// Проверяем на часы
	if strings.Contains(stand, "ч") {
		hoursStr := strings.TrimSuffix(stand, "ч")
		var hours int
		_, err := fmt.Sscanf(hoursStr, "%d", &hours)
		if err == nil {
			return time.Duration(hours) * time.Hour, nil
		}
	}

	// Проверяем на минуты
	if strings.Contains(stand, "мин") {
		minutesStr := strings.TrimSuffix(stand, "мин")
		var minutes int
		_, err := fmt.Sscanf(minutesStr, "%d", &minutes)
		if err == nil {
			return time.Duration(minutes) * time.Minute, nil
		}
	}

	return 0, fmt.Errorf("cannot parse stand duration: %s", stand)
}

func ParseTime(timeStr string, date time.Time) (time.Time, error) {
	// Normalize time string - add leading zero if needed
	if len(timeStr) == 4 && timeStr[1] == ':' { // like "1:10"
		timeStr = "0" + timeStr // becomes "01:10"
	}
	
	var hour, minute int
	_, err := fmt.Sscanf(timeStr, "%d:%d", &hour, &minute)
	if err != nil {
		return time.Time{}, err
	}

	return time.Date(date.Year(), date.Month(), date.Day(), hour, minute, 0, 0, date.Location()), nil
}

// ConvertToTimezone конвертирует время в указанный часовой пояс
func ConvertToTimezone(t time.Time, timezone string) (time.Time, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}
	return t.In(loc), nil
}

// GetTimezoneDifference получает разницу во времени между двумя часовыми поясами
func GetTimezoneDifference(tz1, tz2 string) (time.Duration, error) {
	now := time.Now()

	loc1, err := time.LoadLocation(tz1)
	if err != nil {
		return 0, err
	}

	loc2, err := time.LoadLocation(tz2)
	if err != nil {
		return 0, err
	}

	t1 := now.In(loc1)
	t2 := now.In(loc2)

	_, offset1 := t1.Zone()
	_, offset2 := t2.Zone()

	diff := time.Duration(offset2-offset1) * time.Second

	return diff, nil
}

// FormatDuration форматирует длительность в читаемый вид
func FormatDuration(d time.Duration) string {
	if d < 0 {
		return fmt.Sprintf("-%s", FormatDuration(-d))
	}

	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60

	if hours > 0 {
		return fmt.Sprintf("%dч %dмин", hours, minutes)
	}
	return fmt.Sprintf("%dмин", minutes)
}

// EnhancedTimeConversion улучшенная конвертация времени
func EnhancedTimeConversion(t time.Time, fromTZ, toTZ string) (time.Time, error) {
	
	toLoc, err := time.LoadLocation(toTZ)
	if err != nil {
		return time.Time{}, err
	}
	
	// Конвертируем через UTC для избежания ошибок
	utcTime := t.In(time.UTC)
	return utcTime.In(toLoc), nil
}

// CalculateExactTimeDifference точный расчёт разницы во времени
func CalculateExactTimeDifference(tz1, tz2 string, referenceTime time.Time) (time.Duration, error) {
	loc1, err := time.LoadLocation(tz1)
	if err != nil {
		return 0, err
	}
	
	loc2, err := time.LoadLocation(tz2)
	if err != nil {
		return 0, err
	}
	
	t1 := referenceTime.In(loc1)
	t2 := referenceTime.In(loc2)
	
	return t2.Sub(t1), nil
}

// FormatTimeWithTimezone форматирует время с указанием часового пояса
func FormatTimeWithTimezone(t time.Time, timezone string) (string, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return "", err
	}
	
	localTime := t.In(loc)
	return localTime.Format("15:04 02.01.2006 MST"), nil
}