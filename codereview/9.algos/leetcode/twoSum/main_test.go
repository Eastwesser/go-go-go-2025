package main

import (
    "testing"
)

func BenchmarkTwoSum(b *testing.B) {
    testCases := []int{2, 5, 9}
    bSlice := []int{3, 6, 10}

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
       TwoSum(testCases, bSlice)
    }
}

// go test -bench=. -benchmem