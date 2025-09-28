package main

import (
	"reflect"
	"testing"
)

func TestMergeSortedArrays(t *testing.T) {
	tests := []struct {
		name     string
		arr1     []int
		arr2     []int
		expected []int
	}{
		{
			name:     "оба массива не пустые",
			arr1:     []int{1, 3, 5, 7},
			arr2:     []int{2, 4, 6, 8},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:     "первый массив пустой",
			arr1:     []int{},
			arr2:     []int{2, 4, 6},
			expected: []int{2, 4, 6},
		},
		{
			name:     "второй массив пустой",
			arr1:     []int{1, 3, 5},
			arr2:     []int{},
			expected: []int{1, 3, 5},
		},
		{
			name:     "оба массива пустые",
			arr1:     []int{},
			arr2:     []int{},
			expected: []int{},
		},
		{
			name:     "массивы разной длины",
			arr1:     []int{1, 5, 9},
			arr2:     []int{2, 3, 4, 6, 7, 8},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "отрицательные числа",
			arr1:     []int{-5, -2, 0},
			arr2:     []int{-3, 1, 4},
			expected: []int{-5, -3, -2, 0, 1, 4},
		},
		{
			name:     "дубликаты в массивах",
			arr1:     []int{1, 3, 3, 5},
			arr2:     []int{2, 3, 4},
			expected: []int{1, 2, 3, 3, 3, 4, 5},
		},
		{
			name:     "один элемент в каждом",
			arr1:     []int{2},
			arr2:     []int{1},
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeSortedArrays(tt.arr1, tt.arr2)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("mergeSortedArrays() = %v, ожидается %v", result, tt.expected)
			}

			originalArr1 := make([]int, len(tt.arr1))
			copy(originalArr1, tt.arr1)
			originalArr2 := make([]int, len(tt.arr2))
			copy(originalArr2, tt.arr2)
		})
	}
}
