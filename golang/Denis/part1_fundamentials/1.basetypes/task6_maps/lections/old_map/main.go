package main

import (
	"fmt"
	"sync"
	"unsafe"
)

func main() {
	fmt.Println("=== ПОЛНЫЙ ГИД ПО MAPS В GO ===")
	fmt.Println()

	// 1. ОСНОВЫ СОЗДАНИЯ И РАБОТЫ
	fmt.Println("1. ОСНОВЫ СОЗДАНИЯ И РАБОТЫ")

	// Способ 1: make()
	m1 := make(map[string]int)
	m1["apple"] = 5
	m1["banana"] = 3
	fmt.Printf("m1: %v\n", m1)

	// Способ 2: literal
	m2 := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 8,
	}
	fmt.Printf("m2: %v\n", m2)

	// Способ 3: Пустая мапа
	m3 := map[string]int{}
	fmt.Printf("m3: %v\n", m3)
	fmt.Println()

	// 2. ОПАСНОСТЬ NIL MAP
	fmt.Println("2. ОПАСНОСТЬ NIL MAP")
	var nilMap map[string]int
	fmt.Printf("nilMap == nil: %t\n", nilMap == nil)

	// nilMap["key"] = 42 // PANIC: assignment to entry in nil map
	// value := nilMap["key"] // Не паникует, но бесполезно

	// Правильная инициализация
	if nilMap == nil {
		nilMap = make(map[string]int)
	}
	nilMap["safe"] = 100
	fmt.Printf("После инициализации: %v\n", nilMap)
	fmt.Println()

	// 3. ПРОВЕРКА СУЩЕСТВОВАНИЯ КЛЮЧА
	fmt.Println("3. ПРОВЕРКА СУЩЕСТВОВАНИЯ КЛЮЧА")

	// Неправильно (не отличает отсутствие ключа от нулевого значения)
	value := m2["unknown"]
	fmt.Printf("m2['unknown'] = %d (неясно: нет ключа или значение 0?)\n", value)

	// Правильно
	value, ok := m2["unknown"]
	fmt.Printf("Значение: %d, Ключ существует: %t\n", value, ok)

	value, ok = m2["apple"]
	fmt.Printf("Значение: %d, Ключ существует: %t\n", value, ok)
	fmt.Println()

	// 4. УДАЛЕНИЕ ЭЛЕМЕНТОВ
	fmt.Println("4. УДАЛЕНИЕ ЭЛЕМЕНТОВ")
	fmt.Printf("До удаления: %v\n", m2)
	delete(m2, "banana")
	fmt.Printf("После delete(m2, 'banana'): %v\n", m2)

	// Безопасное удаление (если ключ может не существовать)
	if _, ok := m2["nonexistent"]; ok {
		delete(m2, "nonexistent")
	}
	fmt.Println()

	// 5. ИТЕРАЦИЯ ПО MAP (ПОРЯДОК НЕ ГАРАНТИРУЕТСЯ!)
	fmt.Println("5. ИТЕРАЦИЯ ПО MAP")
	fmt.Println("   Порядок перебора случаен и может меняться между запусками!")

	for key, value := range m2 {
		fmt.Printf("   %s: %d\n", key, value)
	}
	fmt.Println()

	// 6. ЧТО МОЖЕТ БЫТЬ КЛЮЧОМ?
	fmt.Println("6. ТИПЫ КЛЮЧЕЙ")
	// Можно: string, int, float, bool, массивы, структуры с comparable полями
	// Нельзя: slice, map, function

	// Мапы с разными типами ключей
	intKeyMap := map[int]string{1: "one", 2: "two"}
	boolKeyMap := map[bool]string{true: "yes", false: "no"}
	arrayKeyMap := map[[2]int]string{{1, 2}: "coordinates"}

	// Структура как ключ (все поля должны быть comparable)
	type Point struct {
		X, Y int
	}
	structKeyMap := map[Point]string{
		{1, 2}: "point A",
		{3, 4}: "point B",
	}

	fmt.Printf("int keys: %v\n", intKeyMap)
	fmt.Printf("bool keys: %v\n", boolKeyMap)
	fmt.Printf("array keys: %v\n", arrayKeyMap)
	fmt.Printf("struct keys: %v\n", structKeyMap)
	fmt.Println()

	// 7. СРАВНЕНИЕ MAP
	fmt.Println("7. СРАВНЕНИЕ MAP")
	// map1 == map2 // Ошибка: map can only be compared to nil

	// Правильное сравнение
	mapA := map[string]int{"a": 1, "b": 2}
	mapB := map[string]int{"a": 1, "b": 2}
	fmt.Printf("mapA == mapB: %t (нужно сравнивать вручную)\n", mapsEqual(mapA, mapB))

	mapC := map[string]int{"a": 1, "b": 3}
	fmt.Printf("mapA == mapC: %t\n", mapsEqual(mapA, mapC))
	fmt.Println()

	// 8. ПРОИЗВОДИТЕЛЬНОСТЬ И СЛОЖНОСТЬ
	fmt.Println("8. СЛОЖНОСТЬ ОПЕРАЦИЙ")
	// В среднем: O(1) для Get, Put, Delete
	// В худшем случае: O(n) из-за коллизий

	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		m[i] = i * 2
	}

	fmt.Printf("Размер мапы: %d элементов\n", len(m))
	fmt.Printf("Размер структуры мапы: %d байт\n", unsafe.Sizeof(m))
	fmt.Println()

	// 9. SYNC.MAP ДЛЯ КОНКУРЕНТНОСТИ
	fmt.Println("9. SYNC.MAP ДЛЯ КОНКУРЕНТНОСТИ")
	var syncMap sync.Map

	// Store (аналог m[key] = value)
	syncMap.Store("key1", "value1")
	syncMap.Store("key2", "value2")

	// Load (аналог value, ok := m[key])
	if value, ok := syncMap.Load("key1"); ok {
		fmt.Printf("syncMap.Load('key1'): %v\n", value)
	}

	// Delete
	syncMap.Delete("key2")

	// Range (итерация)
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Printf("   %v: %v\n", key, value)
		return true // продолжить итерацию
	})
	fmt.Println()

	// 10. ВНУТРЕННЕЕ УСТРОЙСТВО MAP
	fmt.Println("10. ВНУТРЕННЕЕ УСТРОЙСТВО")
	fmt.Println("   Реализация: хеш-таблица с бакетами")
	fmt.Println("   Каждый бакет содержит 8 пар ключ-значение")
	fmt.Println("   При коллизиях: overflow buckets (цепочка)")
	fmt.Println("   При росте: рехеширование и увеличение бакетов")
	fmt.Println("   Load factor > 6.5: автоматическое увеличение")
	fmt.Println()

	// 11. ЭВАКУАЦИЯ ДАННЫХ
	fmt.Println("11. ЭВАКУАЦИЯ ДАННЫХ")
	fmt.Println("   При рехешировании данные 'эвакуируются' в новые бакеты")
	fmt.Println("   Происходит постепенно (incremental rehashing)")
	fmt.Println("   Старые и новые бакеты сосуществуют во время эвакуации")
	fmt.Println()

	// 12. ПРАКТИЧЕСКИЕ СОВЕТЫ
	fmt.Println("12. ПРАКТИЧЕСКИЕ СОВЕТЫ")
	fmt.Println("   ✅ Всегда инициализируйте мапу через make() или literal")
	fmt.Println("   ✅ Проверяйте существование ключа через second return value")
	fmt.Println("   ✅ Используйте sync.Map для конкурентного доступа")
	fmt.Println("   ❌ Не полагайтесь на порядок итерации")
	fmt.Println("   ❌ Не используйте несравнимые типы как ключи")
	fmt.Println("   ❌ Не работайте с nil map")

fmt.Println("=== ПОТОКОНЕБЕЗОПАСНОСТЬ MAP и СИНХРОНИЗАЦИЯ ===")
	fmt.Println()

	// 1. ДЕМОНСТРАЦИЯ ПОТОКОНЕБЕЗОПАСНОСТИ
	fmt.Println("1. ДЕМОНСТРАЦИЯ ПОТОКОНЕБЕЗОПАСНОСТИ")
	
	unsafeMap := make(map[int]int)
	var wg sync.WaitGroup

	// Конкурентная запись - приведет к data race!
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			unsafeMap[i] = i * 2 // ОПАСНО! Data race!
		}(i)
	}

	wg.Wait()
	fmt.Printf("Размер unsafeMap: %d (может быть меньше 100!)\n", len(unsafeMap))
	fmt.Println()

	// 2. СПОСОБЫ СИНХРОНИЗАЦИИ
	fmt.Println("2. СПОСОБЫ СИНХРОНИЗАЦИИ")

	// Способ 1: sync.Mutex
	fmt.Println("   Способ 1: sync.Mutex")
	safeMap := struct {
		sync.RWMutex
		m map[int]int
	}{m: make(map[int]int)}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.Lock()
			safeMap.m[i] = i * 2
			safeMap.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Printf("   Размер safeMap: %d\n", len(safeMap.m))

	// Чтение с RLock (множественное чтение)
	safeMap.RLock()
	value1 := safeMap.m[42]
	safeMap.RUnlock()
	fmt.Printf("   Значение safeMap[42]: %d\n", value1)
	fmt.Println()

	// Способ 2: sync.Map (специализированная потокобезопасная мапа)
	fmt.Println("   Способ 2: sync.Map")
	var syncMap1 sync.Map

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			syncMap1.Store(i, i*2)
		}(i)
	}
	wg.Wait()

	// Подсчет элементов в sync.Map
	count := 0
	syncMap.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	fmt.Printf("   Размер syncMap: %d\n", count)

	// Получение значения
	if val, ok := syncMap.Load(42); ok {
		fmt.Printf("   Значение syncMap[42]: %d\n", val)
	}
	fmt.Println()

	// 3. КОГДА ИСПОЛЬЗОВАТЬ sync.Map vs sync.Mutex?
	fmt.Println("3. ВЫБОР МЕЖДУ sync.Map И sync.Mutex")
	fmt.Println("   Используйте sync.Map когда:")
	fmt.Println("   - Много конкурентных чтений")
	fmt.Println("   - Мало конкурентных записей")
	fmt.Println("   - Неизвестный набор ключей")
	fmt.Println()
	fmt.Println("   Используйте sync.Mutex когда:")
	fmt.Println("   - Много конкурентных записей")
	fmt.Println("   - Частые обновления")
	fmt.Println("   - Нужен контроль над блокировками")
	fmt.Println()

	// 4. MAKE() VS NEW() - В ЧЕМ РАЗНИЦА?
	fmt.Println("4. MAKE() VS NEW()")

	// new() - выделяет память, возвращает указатель, ИНИЦИАЛИЗИРУЕТ НУЛЯМИ
	ptr := new(map[string]int)
	fmt.Printf("new(map[string]int): %v, тип: %T\n", *ptr, ptr)
	// (*ptr)["key"] = 1 // PANIC: nil map!

	// make() - выделяет память, ИНИЦИАЛИЗИРУЕТ, возвращает значение
	makedMap := make(map[string]int)
	makedMap["key"] = 1
	fmt.Printf("make(map[string]int): %v, тип: %T\n", makedMap, makedMap)

	// Разница на примере slice
	newSlice := new([]int)
	fmt.Printf("new([]int): len=%d, cap=%d, nil=%t\n", 
		len(*newSlice), cap(*newSlice), *newSlice == nil)

	makeSlice := make([]int, 0, 10)
	fmt.Printf("make([]int, 0, 10): len=%d, cap=%d, nil=%t\n", 
		len(makeSlice), cap(makeSlice), makeSlice == nil)
	fmt.Println()

	// 5. ДРУГИЕ ПОЛЕЗНЫЕ ФУНКЦИИ ДЛЯ MAP
	fmt.Println("5. ПОЛЕЗНЫЕ ФУНКЦИИ")

	// Клонирование мапы
	original := map[string]int{"a": 1, "b": 2, "c": 3}
	clone := make(map[string]int)
	for k, v := range original {
		clone[k] = v
	}
	fmt.Printf("Клон мапы: %v\n", clone)

	// Очистка мапы
	for k := range clone {
		delete(clone, k)
	}
	fmt.Printf("После очистки: %v\n", clone)

	// Получение всех ключей
	keys := make([]string, 0, len(original))
	for k := range original {
		keys = append(keys, k)
	}
	fmt.Printf("Ключи мапы: %v\n", keys)

	// Получение всех значений
	values := make([]int, 0, len(original))
	for _, v := range original {
		values = append(values, v)
	}
	fmt.Printf("Значения мапы: %v\n", values)
	fmt.Println()

	// 6. ПАТТЕРНЫ И АНТИПАТТЕРНЫ
	fmt.Println("6. ПАТТЕРНЫ И АНТИПАТТЕРНЫ")

	// Антипаттерн: частые блокировки в цикле
	antiPatternMap := struct {
		sync.Mutex
		m map[int]int
	}{m: make(map[int]int)}

	// ПЛОХО: Блокировка на каждой итерации
	for i := 0; i < 100; i++ {
		antiPatternMap.Lock() // Слишком много блокировок!
		antiPatternMap.m[i] = i
		antiPatternMap.Unlock()
	}

	// ХОРОШО: Одна блокировка на весь блок
	antiPatternMap.Lock()
	for i := 0; i < 100; i++ {
		antiPatternMap.m[i] = i
	}
	antiPatternMap.Unlock()
	fmt.Println("   Избегайте излишних блокировок в циклах!")
	fmt.Println()

	// 7. ТЕСТИРОВАНИЕ НА DATA RACE
	fmt.Println("7. ТЕСТИРОВАНИЕ НА DATA RACE")
	fmt.Println("   Запускайте с флагами:")
	fmt.Println("   go run -race main.go")
	fmt.Println("   go test -race ./...")
	fmt.Println("   go build -race")
	fmt.Println()

	// 8. БЕНЧМАРКИ ПРОИЗВОДИТЕЛЬНОСТИ
	fmt.Println("8. ПРОИЗВОДИТЕЛЬНОСТЬ")
	fmt.Println("   sync.Mutex быстрее для частых записей")
	fmt.Println("   sync.Map быстрее для частых чтений")
	fmt.Println("   Обычная мапа + мьютекс = больше контроля")
	fmt.Println("   sync.Map = проще в использовании")
}

// 9. КАСТОМНАЯ ПОТОКОБЕЗОПАСНАЯ MAP
type SafeMap[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{m: make(map[K]V)}
}

func (sm *SafeMap[K, V]) Set(key K, value V) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

func (sm *SafeMap[K, V]) Get(key K) (V, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, exists := sm.m[key]
	return value, exists
}

func (sm *SafeMap[K, V]) Delete(key K) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.m, key)
}

// Функция для сравнения двух мап
func mapsEqual[K comparable, V comparable](m1, m2 map[K]V) bool {
	if len(m1) != len(m2) {
		return false
	}

	for key, value1 := range m1 {
		if value2, ok := m2[key]; !ok || value1 != value2 {
			return false
		}
	}

	return true
}

/*
	1. Что такое Map?
		Хеш-таблица для пар ключ-значение. Реализована как массив бакетов.
	
	2. Что может быть ключом?
		Любой comparable тип: string, числа, bool, массивы, структуры с comparable полями.
		Нельзя: слайсы, мапы, функции.
	
	3. Почему нет гарантии порядка?
		Намеренно! Чтобы разработчики не полагались на порядок. Порядок зависит от:
			- Хеш-функции
			- Размера таблицы
			- История добавления элементов

	4. Что такое Bucket?
		"Ведро" содержащее до 8 пар ключ-значение. При коллизиях создаются overflow buckets.

	5. Что такое экстра bucket?
		Дополнительные бакеты для разрешения коллизий (цепочка).

	6. Что такое эвакуация данных?
		Процесс перемещения данных в новые бакеты при увеличении размера мапы.
	
	7. Что такое коллизии?
		Когда разные ключи имеют одинаковый хеш. Разрешаются через цепочки бакетов.

	8. Когда использовать sync.Map?
		- Много горутин пишут/читают
		- Rarely written, frequently read
		- Большое количество ключей

	9. Сложность операций
		В среднем: O(1)
		В худшем случае: O(n) (все ключи в одном бакете)

	10. Как устроена мапа?

	type hmap struct {
		count     int    // количество элементов
		buckets   unsafe.Pointer // массив бакетов  
		oldbuckets unsafe.Pointer // старые бакеты при рехеше
		// ... другие поля
	}

	🚨 ОСНОВНЫЕ ОПАСНОСТИ:

	1. Nil Map Panic
	var m map[string]int
	m["key"] = 42 // PANIC

	2. Конкурентный доступ
	// Без синхронизации - data race!
	go func() { m["key"]++ }()
	go func() { m["key"]++ }()

	3. Случайный порядок

	Не полагайтесь на порядок элементов при итерации!

	1. Потоконебезопасность мапы
	// ОПАСНО - data race!
	go func() { m[key] = value }()
	go func() { delete(m, key) }()

	2. Способы синхронизации
		sync.Mutex - полный контроль, подходит для частых записей
		sync.RWMutex - множественное чтение, эксклюзивная запись
		sync.Map - оптимизирован для read-heavy workloads

	3. Разница make() vs new()
	// new() - только выделяет память, возвращает указатель
	ptr := new(Map) // *ptr == nil

	// make() - выделяет и инициализирует, возвращает значение
	m := make(Map) // m != nil

	4. Паттерны использования
		Избегайте излишних блокировок в циклах
		Используйте sync.Map для read-heavy
		Используйте sync.Mutex для write-heavy

	5. Тестирование на data race
		go run -race main.go
		go test -race ./...
*/
