package main

import "fmt"

func mergeSortedArrays(arr1, arr2 []int) []int {
	n := len(arr1)
	m := len(arr2)
	result := make([]int, 0, n+m)

	i, j := 0, 0
	for i < n && j < m {
		if arr1[i] <= arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	for i < n {
		result = append(result, arr1[i])
		i++
	}

	for j < m {
		result = append(result, arr2[j])
		j++
	}

	return result
}

func main() {
	examples := []struct {
		arr1 []int
		arr2 []int
	}{
		{[]int{1, 3, 5, 7}, []int{2, 4, 6, 8}},
		{[]int{1, 2, 3}, []int{}},
		{[]int{}, []int{4, 5, 6}},
		{[]int{1, 5, 9}, []int{2, 3, 4, 6, 7, 8}},
		{[]int{-5, -2, 0}, []int{-3, 1, 4}},
		{[]int{1, 3, 3, 5}, []int{2, 3, 4}},
	}

	for _, example := range examples {
		merged := mergeSortedArrays(example.arr1, example.arr2)
		fmt.Printf("arr1: %v\n", example.arr1)
		fmt.Printf("arr2: %v\n", example.arr2)
		fmt.Printf("Результат: %v\n\n", merged)
	}
}
