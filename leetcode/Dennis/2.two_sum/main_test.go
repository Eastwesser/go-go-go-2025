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

// Тест для ошибочной функции FindTwoSum
func TestFindTwoSum(t *testing.T) {
    tests := []struct {
        name string
        args []int
        want int
    }{
        {"SortedSlice", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11}, // Ожидаем 11, а не 19!
        {"Uneven Numbers", []int{1, 3, 5, 7, 9}, 10}, // 1+9=10
        {"Even Numbers", []int{2, 4, 6, 8, 10}, 12},  // 2+10=12
        {"ShortSlice", []int{5, 10}, 15},              // 5+10=15
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := FindTwoSum(tt.args); got != tt.want {
                t.Errorf("FindTwoSum() = %v, want %v", got, tt.want)
            }
        })
    }
}

// Тест для FindMaxTwoSum
func TestFindMaxTwoSum(t *testing.T) {
    tests := []struct {
        name string
        args []int
        want int
    }{
        {"SortedSlice", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 19},
        {"UnsortedSlice", []int{10, 3, 5, 1, 8, 2, 9, 4, 7, 6}, 19},
        {"Uneven Numbers", []int{1, 3, 5, 7, 9, 11, 13, 15}, 28},
        {"Even Numbers", []int{2, 4, 6, 8, 10}, 18},
        {"NegativeNumbers", []int{-5, -2, -10, -1}, -3}, // -1 + -2 = -3
        {"MixedNumbers", []int{5, -3, 10, -1, 8}, 18},   // 10 + 8 = 18
        {"ShortSlice", []int{5, 10}, 15},
        {"TooShort", []int{5}, 0},
        {"Empty", []int{}, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := FindMaxTwoSum(tt.args); got != tt.want {
                t.Errorf("FindMaxTwoSum() = %v, want %v", got, tt.want)
            }
        })
    }
}

// Тест для FindTwoSumMax (метод двух указателей)
func TestFindTwoSumMax(t *testing.T) {
    tests := []struct {
        name string
        args []int
        want int
    }{
        {"SortedSlice", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 19},
        {"ReverseSorted", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 19},
        {"Uneven Numbers", []int{1, 3, 5, 7, 9, 11, 13, 15}, 28},
        {"Even Numbers", []int{2, 4, 6, 8, 10}, 18},
        {"ShortSlice", []int{5, 10}, 15},
        {"TooShort", []int{5}, 0},
        {"Empty", []int{}, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Для метода двух указателей массив должен быть отсортирован
            sorted := make([]int, len(tt.args))
            copy(sorted, tt.args)
            sort.Ints(sorted)
            
            if got := FindTwoSumMax(sorted); got != tt.want {
                t.Errorf("FindTwoSumMax() = %v, want %v", got, tt.want)
            }
        })
    }
}

// Тест для TwoSumFindPair
func TestTwoSumFindPairSpecific(t *testing.T) {
    tests := []struct {
        name    string
        nums    []int
        target  int
        wantA   int
        wantB   int
        found   bool
    }{
        // Для двух указателей в отсортированном массиве пара будет определяться логикой алгоритма
        {"FoundPair", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11, 1, 10, true},
        // Алгоритм двух указателей найдет 1+8=9 (движение с краев к центру)
        {"FoundMiddle", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, 1, 8, true},
        {"NotFound", []int{1, 2, 3, 4, 5}, 20, 0, 0, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            sorted := make([]int, len(tt.nums))
            copy(sorted, tt.nums)
            sort.Ints(sorted)
            
            a, b, found := TwoSumFindPair(sorted, tt.target)
            
            if found != tt.found {
                t.Errorf("TwoSumFindPair() found = %v, want %v", found, tt.found)
                return
            }
            
            if found && (a != tt.wantA || b != tt.wantB) {
                t.Errorf("TwoSumFindPair() = %v, %v, want %v, %v", a, b, tt.wantA, tt.wantB)
            }
        })
    }
}


// Тест для TwoSumHashMap
func TestTwoSumHashMapSpecific(t *testing.T) {
    tests := []struct {
        name    string
        nums    []int
        target  int
        wantA   int
        wantB   int
        found   bool
    }{
        // Алгоритм найдет первую пару, где complement уже есть в мапе
        // Для [1,2,3,4,5,6,7,8,9,10] и target=11: найдет 5,6
        {"FoundPair", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11, 5, 6, true},
        // Для [10,3,5,1,8] и target=13: 
        // 10→complement=3 (3 нет) → добавить 10
        // 3→complement=10 (10 ЕСТЬ!) → вернуть 3,10
        {"FoundUnsorted", []int{10, 3, 5, 1, 8}, 13, 3, 10, true},
        {"NotFound", []int{1, 2, 3, 4, 5}, 20, 0, 0, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            a, b, found := TwoSumHashMap(tt.nums, tt.target)
            
            if found != tt.found {
                t.Errorf("TwoSumHashMap() found = %v, want %v", found, tt.found)
                return
            }
            
            if found && !((a == tt.wantA && b == tt.wantB) || (a == tt.wantB && b == tt.wantA)) {
                t.Errorf("TwoSumHashMap() = %v, %v, want %v, %v", a, b, tt.wantA, tt.wantB)
            }
        })
    }
}

func TestTwoSumHashMapLogic(t *testing.T) {
    tests := []struct {
        name    string
        nums    []int
        target  int
        found   bool
    }{
        {"FoundPair", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11, true},
        {"FoundUnsorted", []int{10, 3, 5, 1, 8}, 13, true},
        {"NotFound", []int{1, 2, 3, 4, 5}, 20, false},
        {"EmptySlice", []int{}, 5, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            a, b, found := TwoSumHashMap(tt.nums, tt.target)
            
            if found != tt.found {
                t.Errorf("TwoSumHashMap() found = %v, want %v", found, tt.found)
                return
            }
            
            if found {
                // Проверяем, что сумма правильная
                if a + b != tt.target {
                    t.Errorf("TwoSumHashMap() = %v + %v = %v, want sum %v", a, b, a+b, tt.target)
                }
                
                // Проверяем, что числа есть в массиве
                if !contains(tt.nums, a) || !contains(tt.nums, b) {
                    t.Errorf("Numbers %v and %v not found in array %v", a, b, tt.nums)
                }
            }
        })
    }
}

// Вспомогательная функция
func contains(arr []int, num int) bool {
    for _, v := range arr {
        if v == num {
            return true
        }
    }
    return false
}

func TestTwoSumHashMapControlled(t *testing.T) {
    // Специальный массив, где алгоритм гарантированно найдет конкретную пару
    tests := []struct {
        name    string
        nums    []int
        target  int
        wantA   int
        wantB   int
        found   bool
    }{
        // Алгоритм найдет 1,10 потому что 10 идет после 1
        {"SpecificPair", []int{1, 10, 2, 3, 4, 5, 6, 7, 8, 9}, 11, 1, 10, true},
        // Алгоритм найдет 3,10 потому что 10 идет до 3
        {"SpecificPair2", []int{10, 3, 5, 1, 8}, 13, 3, 10, true},
        {"NotFound", []int{1, 2, 3, 4, 5}, 20, 0, 0, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            a, b, found := TwoSumHashMap(tt.nums, tt.target)
            
            if found != tt.found {
                t.Errorf("TwoSumHashMap() found = %v, want %v", found, tt.found)
                return
            }
            
            if found && !((a == tt.wantA && b == tt.wantB) || (a == tt.wantB && b == tt.wantA)) {
                t.Errorf("TwoSumHashMap() = %v, %v, want %v, %v", a, b, tt.wantA, tt.wantB)
            }
        })
    }
}

// Тест для краевых случаев
func TestEdgeCases(t *testing.T) {
    t.Run("EmptySlice", func(t *testing.T) {
        if got := FindMaxTwoSum([]int{}); got != 0 {
            t.Errorf("FindMaxTwoSum([]int{}) = %v, want 0", got)
        }
        
        if got := FindTwoSumMax([]int{}); got != 0 {
            t.Errorf("FindTwoSumMax([]int{}) = %v, want 0", got)
        }
    })
    
    t.Run("SingleElement", func(t *testing.T) {
        if got := FindMaxTwoSum([]int{5}); got != 0 {
            t.Errorf("FindMaxTwoSum([5]) = %v, want 0", got)
        }
        
        if got := FindTwoSumMax([]int{5}); got != 0 {
            t.Errorf("FindTwoSumMax([5]) = %v, want 0", got)
        }
    })
    
    t.Run("TwoElements", func(t *testing.T) {
        if got := FindMaxTwoSum([]int{5, 10}); got != 15 {
            t.Errorf("FindMaxTwoSum([5,10]) = %v, want 15", got)
        }
        
        sorted := []int{5, 10}
        if got := FindTwoSumMax(sorted); got != 15 {
            t.Errorf("FindTwoSumMax([5,10]) = %v, want 15", got)
        }
    })
}
