package main

import (
	"strings"
	"testing"
)

func BenchmarkMultiplyWord(b *testing.B) {
	testCases := []string{
		"abobus 2",
		"BoBoBus 3",
		"racecar 5",
	}
	for i := 0; i < b.N; i++ {
		for _, v := range testCases {
			MultiplyWord(v)
		}
	}
}

// go test -bench=. -benchmem

func TestMultiplyWord(t *testing.T) {
	testCases := []struct {
		name string
		args string
		want string
	}{
		{"Test1", "Aboba 2", "Aboba\nAboba\n"},
		{"Test2", "abobus 3", "abobus\nabobus\nabobus\n"},
		{"Test3", "racecar 5", "racecar\nracecar\nracecar\nracecar\nracecar\n"},
		{"Zero repeats", "test 0", ""},
		{"Single repeat", "hello 1", "hello\n"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := MultiplyWord(tc.args); got != tc.want {
				t.Errorf("MultiplyWord(%q) = %q, want %q", tc.args, got, tc.want)
			}
		})
	}
}

func TestMultiplyWord_Errors(t *testing.T) {
	testCases := []struct {
		name string
		args string
	}{
		{"No number", "hello"},
		{"Not a number", "hello abc"},
		{"Empty string", ""},
		{"Multiple spaces", "hello  5"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := MultiplyWord(tc.args)
			if !strings.Contains(result, "Ошибка") {
				t.Errorf("MultiplyWord(%q) = %q, expected error message", tc.args, result)
			}
		})
	}
}

/*
	go test -v      # Все тесты
	go test -cover  # Покрытие кода тестами
*/
