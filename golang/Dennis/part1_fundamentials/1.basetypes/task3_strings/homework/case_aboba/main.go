package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MultiplyWord(line1 string) string {
	// убираем все пробелы
	parts := strings.Split(strings.TrimSpace(line1), " ")
	if len(parts) != 2 {
		return "Ошибка: нужно ввести слово и число через пробел"
	}

	// ищем число
	word := parts[0]
	numberCount, err := strconv.Atoi(parts[1])
	if err != nil {
		return "Ошибка: второе значение должно быть числом"
	}

	if numberCount == 0 {
		return ""
	}

	// Вычисляем точный размер буфера
	bufferSize := (len(word) + len("\n")) * numberCount
	
	var builder strings.Builder
	builder.Grow(bufferSize)

	for i := 0; i < numberCount; i++ {
		builder.WriteString(word)
		builder.WriteByte('\n')
	}

	return builder.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите слово и число: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	result := MultiplyWord(input)
	fmt.Println(result)
}
