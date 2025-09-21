package main

import (
    "fmt"
    "unsafe"
)

func main() {
    fmt.Println("=== 1. Исследуем целочисленные типы (Integer Types) ===")

    // 1.1. Демонстрация основных целочисленных типов и их размеров
    fmt.Println("\n--- Размеры типов в байтах ---")
    fmt.Printf("int:    %d байт\n", unsafe.Sizeof(int(0)))
    fmt.Printf("int8:   %d байт\n", unsafe.Sizeof(int8(0)))
    fmt.Printf("int16:  %d байт\n", unsafe.Sizeof(int16(0)))
    fmt.Printf("int32:  %d байт\n", unsafe.Sizeof(int32(0)))
    fmt.Printf("int64:  %d байт\n", unsafe.Sizeof(int64(0)))
    fmt.Printf("uint:   %d байт\n", unsafe.Sizeof(uint(0)))
    fmt.Printf("uint8:  %d байт\n", unsafe.Sizeof(uint8(0)))
    fmt.Printf("uint16: %d байт\n", unsafe.Sizeof(uint16(0)))
    fmt.Printf("uint32: %d байт\n", unsafe.Sizeof(uint32(0)))
    fmt.Printf("uint64: %d байт\n", unsafe.Sizeof(uint64(0)))

    // 1.2. Псевдонимы: byte (uint8) и rune (int32)
    fmt.Println("\n--- Псевдонимы byte и rune ---")
    var b byte = 'A' // То же самое, что uint8
    var r rune = 'Я'  // То же самое, что int32 (для Unicode)
    fmt.Printf("byte: %c, %T, %d байт\n", b, b, unsafe.Sizeof(b))
    fmt.Printf("rune: %c, %T, %d байт\n", r, r, unsafe.Sizeof(r))

    // 1.3. Переполнение и "заворачивание" (wrap around)
    fmt.Println("\n--- Переполнение ---")
    var u8 uint8 = 255
    fmt.Printf("Начальное значение: %d\n", u8)
    u8++ // 255 + 1 = 256, но uint8 хранит только до 255
    fmt.Printf("После u8++ (255 + 1): %d\n", u8) // Что выведет? Предсказываем: 0

    var i8 int8 = 127
    fmt.Printf("Начальное значение: %d\n", i8)
    i8++ // 127 + 1 = 128, но int8 max = 127
    fmt.Printf("После i8++ (127 + 1): %d\n", i8) // Что выведет? Предсказываем: -128

    // 1.4. Нельзя смешивать типы (нужно явное приведение)
    fmt.Println("\n--- Приведение типов ---")
    var i32 int32 = 10
    var i64 int64 = 20

    // sum := i32 + i64 // Эта строка вызовет ошибку компиляции: mismatched types int32 and int64
    sum := int64(i32) + i64 // Правильно: явное приведение
    fmt.Printf("int32(10) + int64(20) = %d (%T)\n", sum, sum)
}
