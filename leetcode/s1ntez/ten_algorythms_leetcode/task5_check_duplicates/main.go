package main

import (
	"fmt"
	"sort"
)

// HasDuplicatesMapped - работает на ЛЮБЫХ данных O(n)
func HasDuplicatesMapped(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

// HasDuplicatesTwoPointers - работает ТОЛЬКО на ОТСОРТИРОВАННЫХ данных O(n)
func HasDuplicatesTwoPointers(nums []int) bool {
	// Должны быть отсортированы!
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

// HasDuplicatesTwoPointersSorted - с явной сортировкой O(n log n)
func HasDuplicatesTwoPointersSorted(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	
	// СОРТИРУЕМ сначала!
	sort.Ints(nums)
	
	// Теперь проверяем соседние элементы
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	
	return false
}

func main() {
	firstList := []int{12, 24, 44, 44, 123, 1114}
	secondList := []int{2, 2, 3, 3, 4, 4}
	thirdList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	unsortedList := []int{12, 24, 123, 1114, 44, 44}

	fmt.Println("=== HashMap (всегда работает) ===")
	fmt.Println("Sorted with duplicates:", HasDuplicatesMapped(firstList))     // true
	fmt.Println("Unsorted with duplicates:", HasDuplicatesMapped(unsortedList)) // true

	fmt.Println("\n=== Two Pointers (только отсортированные) ===")
	fmt.Println("Sorted with duplicates:", HasDuplicatesTwoPointers(firstList))     // true
	fmt.Println("Unsorted with duplicates:", HasDuplicatesTwoPointers(unsortedList)) // false ❌

	fmt.Println("\n=== Two Pointers with Sorting ===")
	fmt.Println("Unsorted with duplicates:", HasDuplicatesTwoPointersSorted(unsortedList)) // true ✅
}
