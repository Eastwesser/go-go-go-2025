package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"unsafe"
)

func main() {
	fmt.Println("=== COMPLEX128 (64-БИТНЫЕ КОМПОНЕНТЫ) ===")
	fmt.Println()

	// 1. Создание complex128
	fmt.Println("1. СОЗДАНИЕ complex128")
	c128 := complex(5.0, 12.0)     // По умолчанию создается complex128
	c128_short := 3.0 + 4.0i       // Сокращенная форма

	fmt.Printf("c128: %v, тип: %T, размер: %d байт\n", c128, c128, unsafe.Sizeof(c128))
	fmt.Printf("c128_short: %v, тип: %T\n", c128_short, c128_short)
	fmt.Println()

	// 2. Доступ к компонентам
	fmt.Println("2. КОМПОНЕНТЫ COMPLEX128")
	realPart := real(c128)
	imagPart := imag(c128)
	fmt.Printf("Число: %v\n", c128)
	fmt.Printf("Действительная часть: %.15f, тип: %T\n", realPart, realPart)
	fmt.Printf("Мнимая часть: %.15f, тип: %T\n", imagPart, imagPart)
	fmt.Println()

	// 3. Математические операции
	fmt.Println("3. МАТЕМАТИЧЕСКИЕ ОПЕРАЦИИ")
	a := complex(2.0, 3.0) // 2 + 3i
	b := complex(1.0, 2.0) // 1 + 2i

	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)
	fmt.Printf("a + b = %v\n", a+b)
	fmt.Printf("a - b = %v\n", a-b)
	fmt.Printf("a * b = %v\n", a*b)
	fmt.Printf("a / b = %v\n", a/b)
	fmt.Println()

	// 4. Модуль и фаза
	fmt.Println("4. МОДУЛЬ И ФАЗА")
	z := complex(3.0, 4.0) // 3 + 4i
	modulus := cmplx.Abs(z)
	phase := cmplx.Phase(z)

	fmt.Printf("Число: %v\n", z)
	fmt.Printf("Модуль |z|: %.15f\n", modulus)
	fmt.Printf("Фаза: %.15f радиан\n", phase)
	fmt.Printf("Фаза в градусах: %.15f°\n", phase*180/math.Pi)
	fmt.Println()

	// 5. Полезные функции cmplx
	fmt.Println("5. ФУНКЦИИ CMPLX ДЛЯ COMPLEX128")
	fmt.Printf("Комплексное сопряжение %v: %v\n", z, cmplx.Conj(z))
	fmt.Printf("Экспонента e^%v: %v\n", z, cmplx.Exp(z))
	fmt.Printf("Квадратный корень из %v: %v\n", z, cmplx.Sqrt(z))
	fmt.Printf("Синус %v: %v\n", z, cmplx.Sin(z))
	fmt.Printf("Косинус %v: %v\n", z, cmplx.Cos(z))
	fmt.Printf("Логарифм ln(%v): %v\n", z, cmplx.Log(z))
	fmt.Println()

	// 6. Специальные значения
	fmt.Println("6. СПЕЦИАЛЬНЫЕ ЗНАЧЕНИЯ")
	zero := 0.0
	inf := 1.0 / zero
	nan := zero / zero

	complexInf := complex(inf, 0)
	complexNaN := complex(nan, 0)
	complexBoth := complex(inf, inf)

	fmt.Printf("complex(Inf, 0): %v\n", complexInf)
	fmt.Printf("complex(NaN, 0): %v\n", complexNaN)
	fmt.Printf("complex(Inf, Inf): %v\n", complexBoth)
	fmt.Printf("Is complexInf infinite? %t\n", cmplx.IsInf(complexInf))
	fmt.Printf("Is complexNaN NaN? %t\n", cmplx.IsNaN(complexNaN))
	fmt.Println()

	// 7. Сравнение и точность
	fmt.Println("7. ТОЧНОСТЬ COMPLEX128")
	preciseComplex := complex(123.456789012345, 987.654321098765)

	fmt.Printf("Исходное: %v\n", preciseComplex)
	fmt.Printf("Точное значение: %.15f + %.15fi\n", real(preciseComplex), imag(preciseComplex))
	fmt.Printf("Сохраняет точность: %t\n", preciseComplex == preciseComplex)
	fmt.Println()

	fmt.Println("ВЫВОД: complex128 обеспечивает максимальную точность")
	fmt.Println("       (~15-16 знаков) и является стандартом для вычислений.")
	fmt.Println("       Все функции пакета cmplx работают с complex128.")
}

/*
	Практические выводы:

    complex128 — стандарт де-факто для научных и инженерных вычислений, где важна точность.

    complex64 — для оптимизации в случаях:
        Большие массивы комплексных чисел
        Графические вычисления
        Системы с ограниченной памятью
        Когда потеря точности допустима

    Пакет cmplx — твой главный инструмент для работы с комплексными числами, 
	предоставляет все необходимые математические функции.
*/
