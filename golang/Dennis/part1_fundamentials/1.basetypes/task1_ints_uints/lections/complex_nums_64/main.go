package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"unsafe"
)

/*
	Размер: complex64 = 8 байт, complex128 = 16 байт
	Точность: complex64 показывает меньше знаков после запятой
*/

func main() {
	fmt.Println("=== COMPLEX64 (32-БИТНЫЕ КОМПОНЕНТЫ) ===")
	fmt.Println()

	// 1. Создание complex64
	fmt.Println("1. СОЗДАНИЕ complex64")
	c64 := complex(float32(3), float32(4))    // Явное указание float32
	c64_short := complex(float32(5), float32(12)) // 5 + 12i

	fmt.Printf("c64: %v, тип: %T, размер: %d байт\n", c64, c64, unsafe.Sizeof(c64))
	fmt.Printf("c64_short: %v, тип: %T\n", c64_short, c64_short)
	fmt.Println()

	// 2. Доступ к компонентам
	fmt.Println("2. КОМПОНЕНТЫ COMPLEX64")
	realPart := real(c64)
	imagPart := imag(c64)
	fmt.Printf("Число: %v\n", c64)
	fmt.Printf("Действительная часть: %.7f, тип: %T\n", realPart, realPart)
	fmt.Printf("Мнимая часть: %.7f, тип: %T\n", imagPart, imagPart)
	fmt.Println()

	// 3. Математические операции
	fmt.Println("3. МАТЕМАТИЧЕСКИЕ ОПЕРАЦИИ")
	a := complex(float32(2.0), float32(3.0)) // 2 + 3i
	b := complex(float32(1.0), float32(2.0)) // 1 + 2i

	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)
	fmt.Printf("a + b = %v\n", a+b)
	fmt.Printf("a - b = %v\n", a-b)
	fmt.Printf("a * b = %v\n", a*b)
	fmt.Printf("a / b = %v\n", a/b)
	fmt.Println()

	// 4. Модуль и фаза (преобразуем в complex128 для функций cmplx)
	fmt.Println("4. МОДУЛЬ И ФАЗА")
	z := complex(float32(3.0), float32(4.0)) // 3 + 4i
	modulus := cmplx.Abs(complex128(z))      // Функции cmplx работают с complex128
	phase := cmplx.Phase(complex128(z))

	fmt.Printf("Число: %v\n", z)
	fmt.Printf("Модуль |z|: %.6f\n", modulus)
	fmt.Printf("Фаза: %.6f радиан\n", phase)
	fmt.Printf("Фаза в градусах: %.2f°\n", phase*180/math.Pi)
	fmt.Println()

	// 5. Специальные значения
	fmt.Println("5. СПЕЦИАЛЬНЫЕ ЗНАЧЕНИЯ")
	zero := float32(0.0)
	inf := float32(1.0) / zero
	nan := zero / zero

	complexInf := complex(inf, 0)   // Бесконечность
	complexNaN := complex(nan, 0)   // NaN

	fmt.Printf("complex(Inf, 0): %v\n", complexInf)
	fmt.Printf("complex(NaN, 0): %v\n", complexNaN)
	fmt.Printf("Is complexInf infinite? %t\n", cmplx.IsInf(complex128(complexInf)))
	fmt.Printf("Is complexNaN NaN? %t\n", cmplx.IsNaN(complex128(complexNaN)))
	fmt.Println()

	// 6. Точность complex64
	fmt.Println("6. ТОЧНОСТЬ COMPLEX64")
	preciseComplex := complex(123.456789, 987.654321) // Исходное число
	c64_precise := complex64(preciseComplex)          // Преобразование в complex64

	fmt.Printf("Исходное: %v\n", preciseComplex)
	fmt.Printf("complex64: %v\n", c64_precise)
	fmt.Printf("Потеря точности: %.6f\n", preciseComplex-complex128(c64_precise))
	fmt.Println()

	fmt.Println("ВЫВОД: complex64 использует float32, что экономит память")
	fmt.Println("       но приводит к потере точности (~6-7 знаков).")
}
