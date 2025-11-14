package main

import (
	"fmt"
)

func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	counter := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		counter[s1[i]]++
		counter[s2[i]]--
	}

	for _, count := range counter {
		if count != 0 {
			return false
		}
	}

	return true
}

func main() {
	var stringOne, stringTwo string
	
	fmt.Println("Please, enter 1st word:")
	fmt.Scanln(&stringOne)
	fmt.Println("Please, enter 2nd word:")
	fmt.Scanln(&stringTwo)

	fmt.Println(IsAnagram(stringOne, stringTwo))
}
