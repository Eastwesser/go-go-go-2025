package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {
	tests := [][]int{
		{1, 1, 2},
		{1, 1, 2, 2, 2, 3, 4, 4, 5},
		{1, 2, 3},
		{1, 1, 1},
		{},
	}

	for _, nums := range tests {
		arr := make([]int, len(nums))
		copy(arr, nums)

		length := removeDuplicates(arr)
		fmt.Printf("Было: %v -> Стало: %v (длина: %d)\n", nums, arr[:length], length)
	}
}
