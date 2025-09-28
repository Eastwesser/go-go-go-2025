package main

import "fmt"

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		currentSum := nums[left] + nums[right]

		if currentSum == target {
			return []int{left, right}
		} else if currentSum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

func main() {
	tests := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{1, 3, 4, 6, 8}, 10},
		{[]int{1, 2, 3}, 7},
		{[]int{-3, -1, 0, 2, 5}, -1},
	}

	for _, test := range tests {
		result := twoSum(test.nums, test.target)
		fmt.Printf("nums: %v, target: %d -> ", test.nums, test.target)
		if result != nil {
			fmt.Printf("indices: [%d, %d] (values: %d + %d = %d)\n",
				result[0], result[1], test.nums[result[0]], test.nums[result[1]], test.target)
		} else {
			fmt.Println("no pair found")
		}
	}
}
