package main

import "testing"

func BenchmarkTopFrequentElements(b *testing.B) {
    testCases := []struct {
		nums []int
		target int
	}{
		{[]int{1, 2, 3, 4, 5, 5, 7, 8, 9, 10}, 2}, // 5
        {[]int{12, 17, 11, 15, 22, 27, 111, 15, 32, 37, 112, 15}, 3}, // 15
        {[]int{1, 53, 53, 53, 43, 54, 53, 28, 299, 100}, 4}, // 53
        {[]int{1, 2, 3, 4}, 10}, // 0, edge case
	}
   

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        for _, tc := range testCases {
			TopFrequentElements(tc.nums, tc.target)
		}
       
    }
}
/*
BenchmarkTopFrequentElements-4             18444             55009 ns/op            2624 B/op      36 allocs/op
*/
