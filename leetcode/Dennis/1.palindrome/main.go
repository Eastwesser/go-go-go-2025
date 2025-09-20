package main

import (
	"fmt"
	"strings"
	"unicode"
)

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

func main() {
	// fmt.Println("Palindrome")
	niceWord := "racecar"
	badWord := "not a palindrome"
	fmt.Println(IsPalin(niceWord)) // true
	fmt.Println(IsPalin(badWord)) // false
}
