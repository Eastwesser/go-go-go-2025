package main

import "testing"

func BenchmarkFindTwoSum(b *testing.B) {
    // Тестовые данные: массив чисел и целевая сумма
    testCases := []struct {
        nums   []int
        target int
    }{
        {[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11},  // пара: 1+10
        {[]int{2, 7, 11, 15}, 9},                    // пара: 2+7
        {[]int{1, 3, 5, 7, 9}, 10},                  // пара: 3+7
        {[]int{1, 2, 3, 4}, 10},                     // нет пары
    }

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        for _, tc := range testCases {
            FindTwoSum(tc.nums, tc.target)
        }
    }
}

// go test -bench=. -benchmempackage main
