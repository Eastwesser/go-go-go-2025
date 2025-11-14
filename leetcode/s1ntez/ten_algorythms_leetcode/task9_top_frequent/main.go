package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TopFrequentElements(nums []int, k int) []int {
	frequency := make(map[int]int)

	for _, nums:= range nums {
		frequency[nums]++
	}

	buckets := make([][]int, len(nums)+1)

	for num, freq := range frequency {
		buckets[freq] = append(buckets[freq], num)
	}

	res := make([]int, 0, k)

	for i := len(buckets) - 1; i >= 0 && len(res) < k; i-- {

		if len(buckets[i]) > 0 {
			res = append(res, buckets[i]...)
		}
	}

	return res[:k]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("Введите ваши числа через пробел:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	numStrs := strings.Split(input, " ")
	setOfNums := make([]int, 0, len(numStrs))
	
	for _, str := range numStrs {
		if num, err := strconv.Atoi(str); err == nil {
			setOfNums = append(setOfNums, num)
		}
	}

	fmt.Printf("📊 Введенные числа: %v\n", setOfNums) // Отладочный вывод

	fmt.Println("Введите количество топ элементов (k):")
	var k int
	fmt.Scanln(&k)

	result := TopFrequentElements(setOfNums, k)
	fmt.Printf("🎯 Топ-%d самых частых элементов: %v\n", k, result)
	
	// Дополнительная информация
	fmt.Printf("📈 Частоты: ")
	freq := make(map[int]int)
	for _, num := range setOfNums {
		freq[num]++
	}
	fmt.Println(freq)
}
