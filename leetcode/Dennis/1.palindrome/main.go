package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
	Кейс#1: Original approach with strings.Builder
	BenchmarkIsPalin-4	241653	5274 ns/op	120 B/op	7 allocs/op

	241653 итераций выполнено
	5274 ns/op - 5274 наносекунды на операцию (довольно много)
	120 B/op - 120 байт аллокаций памяти на операцию
	7 allocs/op - 7 выделений памяти на операцию (это много!)

	Аллокации:
	1. strings.ToLower(word) - создает новую строку
	2. strings.Builder - буфер для построения строки  
	3. builder.String() - создает новую строку из буфера
	4. []rune(builder.String()) - конвертация строки в срез рун
	5-7. Дополнительные аллокации в циклах

	Плюсы: Простая логика
	Минусы: Много аллокаций, медленная работа (антипаттерн)

	Как оптимизировать?
		- Убрать strings.Builder - он создает аллокации
		- Убрать strings.ToLower() для всей строки - делаем lowercase посимвольно
		- Использовать two pointers вместо создания нового slice рун
		- Минимизировать аллокации памяти - главный источник замедления
*/
func IsPalin(word string) bool {
	var builder strings.Builder

	for _, r := range strings.ToLower(word) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			builder.WriteRune(r)
		}
	}

	runes := []rune(builder.String())

	for i := 0; i < len(runes) / 2; i++ {
		j := len(runes) - i - 1
		if runes[i] != runes[j] {
			return false
		}
	}

	return true
}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

/*
	Кейс#2: Two pointers with pre-lowered string  
	BenchmarkIsPalinUltraOne-4: 1459 ns/op, 32 B/op, 1 allocs/op

	Аллокации:
	1. strings.ToLower(word) - создает новую строку в нижнем регистре

	Плюсы: 
	- Быстрее оригинала в 3.6 раза
	- Всего 1 аллокация вместо 7
	- Работает с байтами (быстрее рун)

	Минусы:
	- Все еще есть аллокация на strings.ToLower()
	- Конвертация byte->rune для каждого символа в isAlphanumeric
*/
func IsPalinUltraOne(word string) bool {
	lowered := strings.ToLower(word)

	left, right := 0, len(lowered)-1
	for left < right {
		if !isAlphanumeric(rune(lowered[left])) {
			left++
			continue
		}
		if !isAlphanumeric(rune(lowered[right])) {
			right--
			continue
		}
		if lowered[left] != lowered[right] {
			return false
		}
		left++
		right--
	}

	return true
}

/*
	Кейс#3: Two pointers with on-the-fly lowercase
	BenchmarkIsPalinUltraTwo-4: 833 ns/op, 0 B/op, 0 allocs/op

	Аллокации: 0 (zero allocations!) Это самый лучший алгос для палиндромов.
	
	Плюсы:
	- Самая быстрая версия (в 6.3 раза быстрее оригинала)
	- Нет аллокаций памяти
	- On-the-fly обработка символов

	Минусы:
	- Двойная конвертация byte->rune для каждого символа
	- unicode.ToLower() вызывается для каждого символа
*/
func IsPalinUltraTwo(word string) bool {
	left, right := 0, len(word)-1

	for left < right {
		leftChar := unicode.ToLower(rune(word[left]))
		rightChar := unicode.ToLower(rune(word[right]))

		if !isAlphanumeric(leftChar) {
			left++
			continue
		}
		if !isAlphanumeric(rightChar) {
			right--
			continue
		}

		if leftChar != rightChar {
			return false
		}
		left++
		right--
	}

	return true
}

/*
	Кейс#4: Ultra-optimized version (попытка максимальной оптимизации)
	BenchmarkIsPalinUltraThree-4: 888.4 ns/op, 0 B/op, 0 allocs/op

	Реальный результат: немного медленнее UltraTwo (888 ns vs 833 ns)
	
	Ожидалось: ~600-700 ns/op, 0 allocs/op
	Получилось: 888 ns/op, 0 allocs/op

	Анализ:
	- Вложенные циклы добавили overhead вместо ускорения
	- Повторные вызовы unicode.ToLower() в циклах не дали ожидаемого выигрыша
	- Иногда простая логика (UltraTwo) оказывается эффективнее сложной оптимизации

	Плюсы:
	- Все еще 0 аллокаций
	- Хорошая читаемость

	Минусы:
	- Не оправдал ожиданий по скорости
	- Сложнее чем UltraTwo, и медленнее

	Вывод: UltraTwo остается оптимальным выбором
*/
func IsPalinUltraThree(word string) bool {
	left, right := 0, len(word)-1

	for left < right {
		for left < right {
			leftChar := unicode.ToLower(rune(word[left]))
			if isAlphanumeric(leftChar) {
				break
			}
			left++
		}

		for left < right {
			rightChar := unicode.ToLower(rune(word[right]))
			if isAlphanumeric(rightChar) {
				break
			}
			right--
		}

		if left < right {
			leftChar := unicode.ToLower(rune(word[left]))
			rightChar := unicode.ToLower(rune(word[right]))
			if leftChar != rightChar {
				return false
			}
			left++
			right--
		}
	}

	return true
}

func main() {
	fmt.Println("Palindrome task")
	
	niceWord := "racecar"
	badWord := "not a palindrome"

	// Кейс#1 strings.Builder и срезом рун (с 7 аллокациями)
	fmt.Println(IsPalin(niceWord)) // true
	fmt.Println(IsPalin(badWord)) // false

	// Кейс#2 два указателя и один цикл + срез рун (1 аллокация)
	fmt.Println(IsPalinUltraOne(niceWord)) // true
	fmt.Println(IsPalinUltraOne(badWord)) // false
	
	// Кейс#3 два указателя и один цикл (0 аллокаций) для каждого символа unicode.ToLower()
	fmt.Println(IsPalinUltraTwo(niceWord)) // true
	fmt.Println(IsPalinUltraTwo(badWord)) // false

	// Кейс#4 два указателя и один цикл (0 аллокаций) 1 раз unicode.ToLower()
	fmt.Println(IsPalinUltraThree(niceWord)) // true
	fmt.Println(IsPalinUltraThree(badWord)) // false
}
