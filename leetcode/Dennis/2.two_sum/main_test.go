package main

import (
	"sort"
	"testing"
)

func BenchmarkFindTwoSum(b *testing.B) {
	testCases := [][]int{
        {1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
        {10, 20, 30, 40},
        {5, 3, 8, 1, 9},
    }
	
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			FindTwoSum(tc)
		}
	}
}
// BenchmarkFindTwoSum-4   	97697251	         23.49 ns/op	       0 B/op	       0 allocs/op
// BenchmarkFindTwoSum-4   	43769694	         23.64 ns/op	       0 B/op	       0 allocs/op
// BenchmarkFindTwoSum-4   	51782607	         22.70 ns/op	       0 B/op	       0 allocs/op

func BenchmarkFindMaxTwoSum(b *testing.B) {
    testCases := [][]int{
        {1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
        {10, 20, 30, 40},
        {5, 3, 8, 1, 9},
    }
    
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        for _, tc := range testCases {
            FindMaxTwoSum(tc)
        }
    }
}
// BenchmarkFindMaxTwoSum-4   	27254655	        39.18 ns/op	       0 B/op	       0 allocs/op
// BenchmarkFindMaxTwoSum-4   	97054156	        51.77 ns/op	       0 B/op	       0 allocs/op
// BenchmarkFindMaxTwoSum-4   	93821770	        51.65 ns/op	       0 B/op	       0 allocs/op
// BenchmarkFindMaxTwoSum-4   	59148459	        67.52 ns/op	       0 B/op	       0 allocs/op

func BenchmarkTwoSum_TwoPointers(b *testing.B) {
	testCases := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 3, 5, 7, 9, 11, 13, 15},
		{2, 4, 6, 8, 10},
	}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			FindTwoSumMax(tc)
		}
	}
}
// BenchmarkTwoSum_TwoPointers-4   	17415277	   62.32 ns/op	       0 B/op	    0 allocs/op
// BenchmarkTwoSum_TwoPointers-4    13873786       120.8 ns/op         0 B/op       0 allocs/op
// BenchmarkTwoSum_TwoPointers-4   	7135371	       145.3 ns/op	       0 B/op	    0 allocs/op

// Бенчмарк включая сортировку
func BenchmarkTwoPointers_WithSort(b *testing.B) {
	unsortedCases := [][]int{
		{10, 5, 3, 8, 1, 9, 2, 7, 4, 6},
		{15, 3, 11, 7, 1, 9, 13, 5},
		{8, 4, 10, 2, 6},
	}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		for _, tc := range unsortedCases {
			// Включаем время сортировки в бенчмарк
			sorted := make([]int, len(tc))
			copy(sorted, tc)
			sort.Ints(sorted)
			FindTwoSumMax(sorted)
		}
	}
}
// BenchmarkTwoPointers_WithSort-4       541950       2588 ns/op         192 B/op       3 allocs/op
// BenchmarkTwoPointers_WithSort-4   	 1000000	  2917 ns/op	     192 B/op	    3 allocs/op
// BenchmarkTwoPointers_WithSort-4   	 314257	      3807 ns/op	     192 B/op	    3 allocs/op

func BenchmarkTwoSum_HashMap(b *testing.B) {
	testCases := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 3, 5, 7, 9, 11, 13, 15},
		{2, 4, 6, 8, 10},
	}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			TwoSumHashMap(tc, 19) // Ищем максимальную возможную сумму
		}
	}
}
// BenchmarkTwoSum_HashMap-4         514528              7382 ns/op             328 B/op           3 allocs/op
// BenchmarkTwoSum_HashMap-4   	  	 185425	             8018 ns/op	            328 B/op	       3 allocs/op
// BenchmarkTwoSum_HashMap-4   	     336789	             7446 ns/op	            328 B/op	       3 allocs/op

// go test -bench=. -benchmem