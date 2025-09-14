package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("=== СЛАЙСЫ (SLICES) В GO ===")
	fmt.Println()

	// 1. Создание слайсов
	fmt.Println("1. СОЗДАНИЕ СЛАЙСОВ")

	// Способ 1: Literal
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice1: %v, тип: %T, длина: %d, емкость: %d\n", 
		slice1, slice1, len(slice1), cap(slice1))

	// Способ 2: make()
	slice2 := make([]int, 5) // длина=5, емкость=5
	fmt.Printf("slice2: %v, len: %d, cap: %d\n", 
		slice2, len(slice2), cap(slice2))

	slice3 := make([]int, 3, 10) // длина=3, емкость=10
	fmt.Printf("slice3: %v, len: %d, cap: %d\n", 
		slice3, len(slice3), cap(slice3))

	// Способ 3: Из массива
	arr := [5]int{10, 20, 30, 40, 50}
	slice4 := arr[1:4] // [20, 30, 40]
	fmt.Printf("slice4: %v (из массива)\n", slice4)
	fmt.Println()

	// 2. Внутреннее устройство слайса
	fmt.Println("2. ВНУТРЕННЕЕ УСТРОЙСТВО")
	fmt.Printf("Размер структуры слайса: %d байт\n", unsafe.Sizeof(slice1))
	fmt.Println("   Слайс - это структура с 3 полями:")
	fmt.Println("   - Указатель на массив")
	fmt.Println("   - Длина (len)")
	fmt.Println("   - Емкость (cap)")
	fmt.Println()

	// 3. Добавление элементов (append)
	fmt.Println("3. ДОБАВЛЕНИЕ ЭЛЕМЕНТОВ")
	slice := make([]int, 0, 3)
	fmt.Printf("Начальный: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))
	
	slice = append(slice, 1)
	fmt.Printf("После append(1): %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))
	
	slice = append(slice, 2, 3)
	fmt.Printf("После append(2,3): %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))
	
	// Автоматическое увеличение емкости
	slice = append(slice, 4)
	fmt.Printf("После append(4): %v, len: %d, cap: %d (емкость увеличилась!)\n", 
		slice, len(slice), cap(slice))
	fmt.Println()

	// 4. Подводный камень №1: Общие данные
	fmt.Println("4. ПОДВОДНЫЙ КАМЕНЬ: ОБЩИЕ ДАННЫЕ")
	original := []int{1, 2, 3, 4, 5}
	sliceA := original[1:4] // [2, 3, 4]
	sliceB := original[2:5] // [3, 4, 5]
	
	fmt.Printf("original: %v\n", original)
	fmt.Printf("sliceA: %v\n", sliceA)
	fmt.Printf("sliceB: %v\n", sliceB)
	
	// Меняем общий элемент
	sliceA[1] = 999
	fmt.Println("После sliceA[1] = 999:")
	fmt.Printf("original: %v (изменился!)\n", original)
	fmt.Printf("sliceA: %v\n", sliceA)
	fmt.Printf("sliceB: %v (тоже изменился!)\n", sliceB)
	fmt.Println()

	// 5. Копирование слайсов
	fmt.Println("5. КОПИРОВАНИЕ СЛАЙСОВ")
	source := []int{1, 2, 3, 4, 5}
	dest := make([]int, len(source))
	copy(dest, source)
	
	fmt.Printf("source: %v\n", source)
	fmt.Printf("dest: %v\n", dest)
	
	// Меняем копию
	dest[0] = 999
	fmt.Println("После dest[0] = 999:")
	fmt.Printf("source: %v (не изменился)\n", source)
	fmt.Printf("dest: %v\n", dest)
	fmt.Println()

	// 6. Подводный камень №2: Утечки памяти
	fmt.Println("6. ПОДВОДНЫЙ КАМЕНЬ: УТЕЧКИ ПАМЯТИ")
	bigSlice := make([]int, 0, 1000)
	for i := 0; i < 10; i++ {
		bigSlice = append(bigSlice, i)
	}
	
	fmt.Printf("После добавления 10 элементов: len=%d, cap=%d\n", 
		len(bigSlice), cap(bigSlice))
	
	// Оставляем только нужные элементы (избегаем утечки)
	smallSlice := make([]int, len(bigSlice))
	copy(smallSlice, bigSlice)
	fmt.Printf("После копирования: len=%d, cap=%d\n", 
		len(smallSlice), cap(smallSlice))
	fmt.Println()

	// 7. Срезы (slicing) и емкость
	fmt.Println("7. СРЕЗЫ И ЕМКОСТЬ")
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice5 := data[2:5]
	fmt.Printf("data: %v\n", data)
	fmt.Printf("data[2:5]: %v, len: %d, cap: %d\n", 
		slice5, len(slice5), cap(slice5))
	
	// Можно расширить слайс в пределах емкости
	extended := slice5[:7]
	fmt.Printf("slice5[:7]: %v, len: %d, cap: %d\n", 
		extended, len(extended), cap(extended))
	fmt.Println()

	// 8. Передача слайсов в функции
	fmt.Println("8. ПЕРЕДАЧА В ФУНКЦИИ")
	numbers := []int{1, 2, 3}
	fmt.Printf("До функции: %v\n", numbers)
	modifySlice(numbers)
	fmt.Printf("После функции: %v (изменился!)\n", numbers)
	fmt.Println()

	// 9. Многомерные слайсы
	fmt.Println("9. МНОГОМЕРНЫЕ СЛАЙСЫ")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		make([]int, 2),
	}
	fmt.Printf("matrix: %v\n", matrix)
	matrix[0][1] = 999
	fmt.Printf("После изменения: %v\n", matrix)
	fmt.Println()

	// 10. Практические рекомендации
	fmt.Println("10. КОГДА ЧТО ИСПОЛЬЗОВАТЬ?")
	fmt.Println("   make([]T, 0, capacity) - когда знаете примерный размер")
	fmt.Println("   append() - для добавления элементов")
	fmt.Println("   copy() - когда нужно избежать общих данных")
	fmt.Println("   len(slice) - всегда проверяйте длину перед доступом!")
	fmt.Println()
	fmt.Println("   ПРЕИМУЩЕСТВА перед массивами:")
	fmt.Println("   + Динамический размер")
	fmt.Println("   + Передаются по ссылке (эффективно)")
	fmt.Println("   + Богатые возможности работы")
}

func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 999 // Изменяет оригинальный слайс!
	}
	fmt.Printf("   В функции: %v\n", s)
}
