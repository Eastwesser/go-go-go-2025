package main

import (
	"fmt"
	"sort"
)

// Неправильная версия с ошибками
func FindTwoSum(nums []int) int {
	left, right := 0, len(nums) - 1
	counter := 0

	for left < right {
		for i := 0; i < len(nums) / 2; i++ {
			j := len(nums) - i - 1

			if nums[i] + nums[j] > counter {
				counter = nums[i] + nums[j]
				return counter
			} 
		}
		left++
		right--
	}

	return counter
}

// FindMaxTwoSum находит максимальную сумму двух чисел в слайсе
func FindMaxTwoSum(nums []int) int {
    if len(nums) < 2 {
        return 0
    }

	// Находим два максимальных числа
    max1, max2 := nums[0], nums[1]
	if max2 > max1 {
		max1, max2 = max2, max1
	}
	
	for i := 2; i < len(nums); i++ {
		if nums[i] > max1 {
			max2 = max1
			max1 = nums[i]
		} else if nums[i] > max2 {
			max2 = nums[i]
		}
	}
		
	return max1 + max2
}

// FindTwoSumMax находит максимальную сумму двух чисел методом двух указателей
// Предполагает, что слайс отсортирован по возрастанию
func FindTwoSumMax(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	
	left, right := 0, len(nums)-1
	maxSum := nums[left] + nums[right]
	
	for left < right {
		currentSum := nums[left] + nums[right]
		if currentSum > maxSum {
			maxSum = currentSum
		}
		
		// Двигаем указатели - в отсортированном массиве максимальная сумма
		// будет достигаться при движении к центру от краев
		if nums[left] < nums[right] {
			left++
		} else {
			right--
		}
	}
	
	return maxSum
}

// TwoSumFindPair находит пару чисел, дающую целевую сумму
// Классическая задача Two Sum методом двух указателей
func TwoSumFindPair(nums []int, target int) (int, int, bool) {
	left, right := 0, len(nums)-1
	
	for left < right {
		sum := nums[left] + nums[right]
		
		if sum == target {
			return nums[left], nums[right], true
		} else if sum < target {
			left++ // Увеличиваем сумму
		} else {
			right-- // Уменьшаем сумму
		}
	}
	
	return 0, 0, false
}

// TwoSumHashMap - реализация через хэш-таблицу для несортированных данных
func TwoSumHashMap(nums []int, target int) (int, int, bool) {
	seen := make(map[int]bool)
	
	for _, num := range nums {
		complement := target - num
		if seen[complement] {
			return complement, num, true
		}
		seen[num] = true
	}
	
	return 0, 0, false
}

func main() {
	fmt.Println("Two Sum")
	// Ошибочная версия
	numList1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 11
	fmt.Println(FindTwoSum(numList1))
	// Правильная версия
	numList2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 19 (9 + 10)
    fmt.Println(FindMaxTwoSum(numList2)) // Output: 19


	fmt.Println("Two Sum - Метод двух указателей")
	// Для метода двух указателей массив должен быть отсортирован!
	sortedNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 1. Находим максимальную сумму двух чисел
	maxSum := FindTwoSumMax(sortedNums)
	fmt.Printf("Максимальная сумма двух чисел: %d\n", maxSum) // 19
	// 2. Находим пару чисел для конкретной цели
	target := 11
	a, b, found := TwoSumFindPair(sortedNums, target)
	if found {
		fmt.Printf("Пара чисел дающая сумму %d: %d и %d\n", target, a, b)
	} else {
		fmt.Printf("Пара для суммы %d не найдена\n", target)
	}
	// 3. Пример с несортированным массивом (сначала сортируем)
	unsortedNums := []int{10, 3, 5, 1, 8, 2, 9, 4, 7, 6}
	sort.Ints(unsortedNums) // Сортируем для использования двух указателей
	fmt.Printf("После сортировки: %v\n", unsortedNums)
	maxSumUnsorted := FindTwoSumMax(unsortedNums)
	fmt.Printf("Максимальная сумма после сортировки: %d\n", maxSumUnsorted)
}
