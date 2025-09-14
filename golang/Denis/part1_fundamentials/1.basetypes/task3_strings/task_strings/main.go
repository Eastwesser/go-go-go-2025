package main

import (
	"fmt"
	"strings"
)

func multiplyWord(line1 string) string {
	var builder strings.Builder


	builtStr := builder.String()
	fmt.Printf("   strings.Builder: длина %d байт\n", len(builtStr))
	fmt.Println()
	strconv.Atoi() "5" -> 5

	for i := 0; i < number; i++ {
		builder.WriteString(line1)
	}

	return builtStr
}

func main() {
	var stringLine string
	fmt.Scan(&stringLine)
	fmt.Println(multiplyWord(stringLine))
}
