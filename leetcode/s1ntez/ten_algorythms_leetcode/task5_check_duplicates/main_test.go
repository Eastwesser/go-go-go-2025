package main

import (
    "testing"
)

func BenchmarkHasDuplicatesMapped(b *testing.B) {
    testCases := [][]int{
        {12, 24, 44, 44, 123, 1114},
        {1, 2, 3, 4, 5},
        {5, 4, 3, 2, 1, 5},
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        for _, tc := range testCases {
            HasDuplicatesMapped(tc)
        }
    }
}

func BenchmarkHasDuplicatesTwoPointersSorted(b *testing.B) {
    testCases := [][]int{
        {12, 24, 44, 44, 123, 1114},
        {1, 2, 3, 4, 5},
        {5, 4, 3, 2, 1, 5},
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        for _, tc := range testCases {
            HasDuplicatesTwoPointersSorted(tc)
        }
    }
}

// go test -bench=. -benchmempackage main

