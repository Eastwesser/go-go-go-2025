package main

import "fmt"

func intersection(nums1, nums2 []int) []int {
	set := make(map[int]bool)
	for _, num := range nums1 {
		set[num] = true
	}

	result := []int{}
	for _, num := range nums2 {
		if set[num] {
			result = append(result, num)
			delete(set, num)
		}
	}

	return result
}

func main() {
	tests := []struct {
		arr1, arr2 []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}},
		{[]int{1, 2, 3}, []int{4, 5, 6}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{1, 2}},
	}

	for _, test := range tests {
		result := intersection(test.arr1, test.arr2)
		fmt.Printf("arr1: %v, arr2: %v -> пересечение: %v\n", test.arr1, test.arr2, result)
	}
}
