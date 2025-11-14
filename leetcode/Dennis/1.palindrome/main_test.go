package main

import "testing"

func BenchmarkIsPalin(b *testing.B) {
	testCases := []string{
		"racecar",
		"A man, a plan, a canal: Panama",
		"not a palindrome",
		"",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			IsPalin(tc)
		}
	}
}

func BenchmarkIsPalinUltraOne(b *testing.B) {
	testCases := []string{
		"racecar",
		"A man, a plan, a canal: Panama",
		"not a palindrome",
		"",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			IsPalinUltraOne(tc)
		}
	}
}

func BenchmarkIsPalinUltraTwo(b *testing.B) {
	testCases := []string{
		"racecar",
		"A man, a plan, a canal: Panama",
		"not a palindrome",
		"",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			IsPalinUltraTwo(tc)
		}
	}
}

func BenchmarkIsPalinUltraThree(b *testing.B) {
	testCases := []string{
		"racecar",
		"A man, a plan, a canal: Panama",
		"not a palindrome",
		"",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			IsPalinUltraThree(tc)
		}
	}
}

// go test -bench=. -benchmem

func TestIsPalin(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{"Normal", "racecar", true},
		{"Camel", "RaDaR", true},
		{"Not palindrome", "dinosaur", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalin(tt.args); got != tt.want {
				t.Errorf("IsPalin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPalinUltraOne(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{"Normal", "racecar", true},
		{"Camel", "RaDaR", true},
		{"Not palindrome", "dinosaur", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalinUltraOne(tt.args); got != tt.want {
				t.Errorf("IsPalinUltraOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPalinUltraTwo(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{"Normal", "racecar", true},
		{"Camel", "RaDaR", true},
		{"Not palindrome", "dinosaur", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalinUltraTwo(tt.args); got != tt.want {
				t.Errorf("IsPalinUltraTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPalinUltraThree(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{"Normal", "racecar", true},
		{"Camel", "RaDaR", true},
		{"Not palindrome", "dinosaur", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalinUltraThree(tt.args); got != tt.want {
				t.Errorf("IsPalinUltraThree() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
	# Все тесты
	go test -v

	# Только benchmark
	go test -bench=. -benchmem -run=^$

	# Benchmark с профилем CPU
	go test -bench=. -cpuprofile=cpu.out

	# Benchmark с профилем памяти
	go test -bench=. -memprofile=mem.out

	# Fuzz тестирование
	go test -fuzz=FuzzIsPalindrome -fuzztime=30s

	# Покрытие кода
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out
*/
