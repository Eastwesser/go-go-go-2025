package main

import (
	"fmt"
	"testing"
)

func BenchmarkQuickSort(b *testing.B) {
    // Генерируем тестовые данные
    generateUsers := func(n int) []User {
        users := make([]User, n)
        for i := 0; i < n; i++ {
            users[i] = User{
                ID:   n - i, // Обратный порядок для худшего случая
                Name: fmt.Sprintf("User%d", n-i),
                Age:  (n - i) % 100,
            }
        }
        return users
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        users := generateUsers(1000) // 1000 пользователей
        QuickSort(users, func(a, b User) bool {
            return a.ID < b.ID
        })
    }
}

// go test -bench=. -benchmem
