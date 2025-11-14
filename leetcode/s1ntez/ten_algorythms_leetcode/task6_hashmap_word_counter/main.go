package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func WordCounter(text string) map[string]int {
	// Более продвинутая очистка слов
	cleanText := strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1 // Удаляем всю пунктуацию
		}
		return unicode.ToLower(r) // Приводим к нижнему регистру
	}, text)

	words := strings.Fields(cleanText)
	seen := make(map[string]int)

	for _, word := range words {
		seen[word]++
	}

	return seen
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your sentence:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Убираем перенос строки
	
	wordCounts := WordCounter(input)
	
	fmt.Println("\n📊 Word Count Results:")
	for word, count := range wordCounts {
		fmt.Printf("'%s': %d\n", word, count)
	}
	
	fmt.Printf("\n📈 Total unique words: %d\n", len(wordCounts))
}
