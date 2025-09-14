package main

import (
	"fmt"
)

// Функция для демонстрации передачи массива
func modifyArray(arr [3]int) {
	arr[0] = 999 // Изменяем копию массива
	fmt.Printf("   В функции: %v\n", arr)
}

func main() {
	// fmt.Println("7. ПЕРЕДАЧА В ФУНКЦИЮ (ВАЖНЫЙ НЮАНС)")
	arr8 := [3]int{1, 2, 3}
	
	fmt.Printf("До функции: %v\n", arr8)

	modifyArray(arr8) // Массив передается ПО ЗНАЧЕНИЮ!
	
	fmt.Printf("После функции: %v (не изменился!)\n", arr8)
}
