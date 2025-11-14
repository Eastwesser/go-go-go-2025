package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	fmt.Println("=== СТРОКИ В GO ===")
	fmt.Println()

	// 1. Базовое создание и свойства
	fmt.Println("1. БАЗОВЫЕ СВОЙСТВА СТРОК")
	str := "Hello, 世界" // Строка с ASCII и Unicode
	fmt.Printf("Строка: %s\n", str)
	fmt.Printf("Длина (len): %d байт\n", len(str)) ///
	fmt.Printf("Длина (utf8.RuneCount): %d символов\n", utf8.RuneCountInString(str))
	fmt.Printf("Размер строки: %d байт\n", unsafe.Sizeof(str))
	fmt.Println()

	// 2. Неизменяемость (Immutable)
	fmt.Println("2. НЕИЗМЕНЯЕМОСТЬ СТРОК")
	// str[0] = 'h' // Ошибка компиляции: cannot assign to str[0]
	fmt.Println("   Строки неизменяемы! Нельзя изменить отдельный символ.")
	fmt.Println()

	// 3. Доступ к символам (байты vs руны)
	fmt.Println("3. ДОСТУП К СИМВОЛАМ")
	fmt.Println("   По индексу получаем БАЙТЫ:")
	for i := 0; i < len(str); i++ {
		fmt.Printf("   str[%d] = %c (байт: %d)\n", i, str[i], str[i])
	}
	fmt.Println()

	fmt.Println("   Для Unicode символов используем RANGE (получаем РУНЫ):")
	for i, r := range str {
		fmt.Printf("   str[%d] = %c (руна: %U)\n", i, r, r)
	}
	fmt.Println()

	// 4. Подстроки (substring)
	fmt.Println("4. ПОДСТРОКИ")
	sub := str[7:10] // Берем байты, а не символы!
	fmt.Printf("   str[7:10] = '%s' (может обрезать символ!)\n", sub)
	fmt.Println()

	// 5. Преобразование в []byte и []rune
	fmt.Println("5. ПРЕОБРАЗОВАНИЕ ТИПОВ")
	bytes := []byte(str)
	runes := []rune(str)
	fmt.Printf("   []byte: %v\n", bytes)
	fmt.Printf("   []rune: %v\n", runes)
	fmt.Printf("   Размер []byte: %d байт\n", unsafe.Sizeof(bytes))
	fmt.Printf("   Размер []rune: %d байт\n", unsafe.Sizeof(runes))
	fmt.Println()

	// 6. Эффективное изменение строк
	fmt.Println("6. ЭФФЕКТИВНОЕ ИЗМЕНЕНИЕ СТРОК")
	
	// Способ 1: Конкатенация (плохо для многих операций)
	result := ""
	for i := 0; i < 5; i++ {
		result += str // Создается новая строка каждый раз!
	}
	fmt.Printf("   Конкатенация: длина %d байт\n", len(result))
	
	// Способ 2: strings.Builder (РЕКОМЕНДУЕТСЯ)
	var builder strings.Builder
	for i := 0; i < 5; i++ {
		builder.WriteString(str)
	}
	builtStr := builder.String()
	fmt.Printf("   strings.Builder: длина %d байт\n", len(builtStr))
	fmt.Println()

	// 7. Сравнение производительности
	fmt.Println("7. КОГДА ЧТО ИСПОЛЬЗОВАТЬ?")
	fmt.Println("   []byte:")
	fmt.Println("   - Когда работаем с бинарными данными")
	fmt.Println("   - Низкоуровневые операции")
	fmt.Println("   - Частые изменения небольших строк")
	fmt.Println()
	fmt.Println("   []rune:")
	fmt.Println("   - Когда нужен посимвольный доступ к Unicode")
	fmt.Println("   - Обработка текста с символами outside BMP")
	fmt.Println()
	fmt.Println("   strings.Builder:")
	fmt.Println("   - Многократная конкатенация строк")
	fmt.Println("   - Построение больших строк")
	fmt.Println()
	fmt.Println("   string:")
	fmt.Println("   - Для хранения и передачи данных")
	fmt.Println("   - Когда данные не нужно изменять")
	fmt.Println()

	// 8. Практические примеры
	fmt.Println("8. ПРАКТИЧЕСКИЕ ПРИМЕРЫ")
	
	// Замена символов через []rune
	runes = []rune(str)
	runes[0] = 'h' // Меняем первый символ
	modifiedStr := string(runes)
	fmt.Printf("   Замена символа: '%s' -> '%s'\n", str, modifiedStr)
	
	// Обрезка строки с учетом Unicode
	properSub := string([]rune(str)[7:9]) // Берем символы, а не байты
	fmt.Printf("   Правильная подстрока: '%s'\n", properSub)
	fmt.Println()

	// 9. Полезные функции пакета strings
	fmt.Println("9. ПОЛЕЗНЫЕ ФУНКЦИИ strings")
	fmt.Printf("   Contains: %t\n", strings.Contains(str, "世界"))
	fmt.Printf("   HasPrefix: %t\n", strings.HasPrefix(str, "Hello"))
	fmt.Printf("   Count: %d\n", strings.Count(str, "l"))
	fmt.Printf("   ToUpper: %s\n", strings.ToUpper(str))
	fmt.Printf("   Replace: %s\n", strings.Replace(str, "Hello", "Hi", 1))
	fmt.Println()

	// 10. Важные нюансы
	fmt.Println("10. ВАЖНЫЕ НЮАНСЫ")
	byteStr := []byte("hello")
	byteStr[0] = 'H' // Можно менять!
	fmt.Printf("   []byte изменяем: %s\n", string(byteStr))
	
	runeStr := []rune("hello")
	runeStr[0] = 'H' // Можно менять!
	fmt.Printf("   []rune изменяем: %s\n", string(runeStr))
	fmt.Println()
	fmt.Println("   ВЫВОД: Для изменения строк преобразуйте в []byte или []rune,")
	fmt.Println("   изменяйте, затем конвертируйте обратно в string.")


	// ПЛОХО - работаем с байтами
	for i := 0; i < len(str); i++ {
		fmt.Print(string(str[i])) // Испортит Unicode символы!
	}

	// ХОРОШО - работаем с рунами
	for _, r := range str {
		fmt.Print(string(r)) // Правильный обход символов
	}

	// ОПАСНО - работа с байтами
	sub1 := str[7:9] // Может обрезать символ!
	fmt.Println(sub1)

	// БЕЗОПАСНО - работа с рунами
	runes1 := []rune(str)
	sub2 := string(runes1[7:9]) // Правильно берем символы
	fmt.Println(sub2)

	// Для ASCII текста (английский)
	dataEng := []byte("hello") // Экономичнее
	dataEng[0] = 'H'

	// Для Unicode текста (китайский, emoji)
	dataNotEng := []rune("hello 世界") // Правильнее
	dataNotEng[6] = '世'
}

/*
	Best Practices:

    Для конкатенации → strings.Builder

    Для изменения символов → []rune

    Для бинарных данных → []byte

    Для подсчета символов → utf8.RuneCountInString()

    Для обхода текста → for _, r := range str
*/
