package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func StringPalindrome(str string) bool {
	left, right := 0, len(str) -1

	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func StringPalindromeRus(rusStr string) bool {
	runes := []rune(rusStr) // конвертация в руны происходит уже внутри функции, а не в аргументах
	left, right := 0, len(runes ) - 1

	for left < right {
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func NumberPalindrome(num int) bool {
	if num < 0 {
		return false
	}
	
	original := num
	reversed := 0

	for num > 0 {
		reversed = reversed * 10 + num % 10 // Формула разворота числа
        num = num / 10                      // Убираем последнюю цифру
	}

	return original == reversed
}

func checkInput(input string) string {
	if _, err := strconv.Atoi(input); err == nil {
		return "number"
	} 
	
	for _, r := range input {
		if unicode.Is(unicode.Cyrillic, r) {
			return "rusWord"
		}
	}

	return "engWord"
}

func main() {
	var newWord string
	fmt.Println("Введите ваше слово (или число), чтобы проверить его на палиндром:")
	fmt.Scanln(&newWord)
	
	inputChoise := checkInput(newWord)

	switch inputChoise {
	case "engWord":
		result := StringPalindrome(newWord)
		fmt.Println(result) // не работает с русскими буквами!!!
	case "rusWord":
		result := StringPalindromeRus(newWord)
		fmt.Println(result) // выдаст "палиндромность" даже на русском, так как мы использовали руны
	case "number":
		number, _ := strconv.Atoi(newWord)
		result := NumberPalindrome(number)
		fmt.Println(result) // выдаст "палиндромность" даже на число!!!
	default:
		fmt.Println("Введите, пожалуйста, слово или число =_=")
	}
	
}
