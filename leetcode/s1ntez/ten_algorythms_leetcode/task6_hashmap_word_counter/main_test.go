package main

import "testing"

func BenchmarkWordCounter(b *testing.B) {
    testCases := []string{
        "Ya pomnu chudnoye mgnovenye",
        "Peredo mnoy yavilas ti", 
        "Da, da, da, eto Pushkin!",
        "Hello, world! This is a test sentence with multiple words.",
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        for _, tc := range testCases {
            WordCounter(tc)
        }
    }
}

// go test -bench=. -benchmempackage main
