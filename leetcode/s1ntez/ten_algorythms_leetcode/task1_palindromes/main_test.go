package main

import (
    "testing"
)

func BenchmarkStringPalindrome(b *testing.B) {
    testCases := []string{       
        "radar",
        "racecar",
        "willow",
    }

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        for _, tc := range testCases {
			StringPalindrome(tc)
		}
       
    }
}

func BenchmarkStringPalindromeRus(b *testing.B) {
    testCases := []string{       
        "радар",
        "комок",
        "якорь",
    }

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        for _, tc := range testCases {
			StringPalindromeRus(tc)
		}
       
    }
}

func BenchmarkNumberPalindrome(b *testing.B) {
    testCases := []int{       
        10101,
        1234321,
        89162730555,
    }

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        for _, tc := range testCases {
			NumberPalindrome(tc)
		}
       
    }
}

// go test -bench=. -benchmempackage main
