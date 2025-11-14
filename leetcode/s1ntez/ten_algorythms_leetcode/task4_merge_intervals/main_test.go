package main

import "testing"

func BenchmarkMergeIntervals(b *testing.B) {
    testCases := []struct {
        name string
        intervals [][]int
    }{
        {"Small", [][]int{{1,3}, {2,6}, {8,10}, {15,18}}},
        {"Touching", [][]int{{1,4}, {4,5}}},
        {"Contained", [][]int{{1,10}, {2,3}, {4,5}, {6,7}}},
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        for _, tc := range testCases {
            // Создаем копию для каждой итерации
            intervalsCopy := make([][]int, len(tc.intervals))
            copy(intervalsCopy, tc.intervals)
            MergeIntervals(intervalsCopy)
        }
    }
}
/*
BenchmarkMergeIntervals-4         185340              9723 ns/op             744 B/op         17 allocs/op
*/
