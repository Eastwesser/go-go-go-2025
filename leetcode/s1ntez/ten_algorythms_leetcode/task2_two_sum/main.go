package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"strconv"
)

func FindTwoSum(nums []int, target int) []int {
	left, right := 0, len(nums) - 1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите числа через пробел:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Парсим числа
	numStrs := strings.Split(input, " ")
	setOfNums := make([]int, 0, len(numStrs))

	// Сортируем для Two Pointers
	sort.Ints(setOfNums)
	fmt.Printf("Отсортированный массив: %v\n", setOfNums)

	fmt.Println("Введите целевую сумму:")
	sumInput, _ := reader.ReadString('\n')
	theSum, _ := strconv.Atoi(strings.TrimSpace(sumInput))

	res := FindTwoSum(setOfNums, theSum)
	if res != nil {
		fmt.Printf("🎯 Найдена пара: nums[%d]=%d + nums[%d]=%d = %d\n", 
			res[0], setOfNums[res[0]], res[1], setOfNums[res[1]], theSum)
	} else {
		fmt.Println("❌ Пара не найдена")
	}
}
