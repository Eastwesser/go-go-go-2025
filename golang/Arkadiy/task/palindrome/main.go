package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	runes := []rune(s)
	left, right := 0, len(runes)-1

	for left < right {
		for left < right && !isAlphanumeric(runes[left]) {
			left++
		}
		for left < right && !isAlphanumeric(runes[right]) {
			right--
		}
		if unicode.ToLower(runes[left]) != unicode.ToLower(runes[right]) {
			return false
		}

		left++
		right--
	}
	return true
}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func main() {
	testCases := []string{
		"ABOBA",
		"Blue",
	}

	for _, test := range testCases {
		result := isPalindrome(test)
		fmt.Printf(test, result)
	}
}
