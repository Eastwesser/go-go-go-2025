package main

import (
    "fmt"
    "math"
    "unsafe"
)

func main() {
    var f64 float64 = 123.456789012345
    fmt.Printf("float64: %.15f, %T, %d байт\n", f64, f64, unsafe.Sizeof(f64))

    // Специальные значения: +Inf, -Inf, NaN
	zero := 0.0
    positiveInf := 1.0 / zero // Положительная бесконечность (if 1.0 / 0.0 -> Compiler says DivByZero)
    negativeInf := -1.0 / zero // Отрицательная бесконечность (Compiler says DivByZero)
    nan := zero / zero // Not a Number (Compiler says DivByZero)

    fmt.Printf("+Inf: %f\n", positiveInf)
    fmt.Printf("-Inf: %f\n", negativeInf)
    fmt.Printf("NaN: %f\n", nan)

	// Важнейшее свойство NaN: он не равен самому себе!
	fmt.Printf("NaN == NaN? %t (Это главный признак для проверки NaN)\n", nan == nan)

    // Проверка на специальные значения (используем math package)
    fmt.Printf("math.IsInf(+Inf, 1): %t\n", math.IsInf(positiveInf, 1))   // Проверка на +бесконечность
	fmt.Printf("math.IsInf(-Inf, -1): %t\n", math.IsInf(negativeInf, -1)) // Проверка на -бесконечность
	fmt.Printf("math.IsInf(+Inf, 0): %t (Любая бесконечность)\n", math.IsInf(positiveInf, 0))
	fmt.Printf("math.IsNaN(NaN): %t\n", math.IsNaN(nan)) // Правильная проверка на NaN

    // Неточность вычислений с плавающей точкой
    f1 := 0.1
    f2 := 0.2
    f3 := 0.3

    fmt.Printf("f1 (0.1) = %.20f\n", f1)
	fmt.Printf("f2 (0.2) = %.20f\n", f2)
	fmt.Printf("f3 (0.3) = %.20f\n", f3)
	
	sum := f1 + f2
	fmt.Printf("f1 + f2 (0.1 + 0.2) = %.20f\n", sum)

	// Прямое сравнение (НЕПРАВИЛЬНО для float)
	fmt.Printf("0.1 + 0.2 == 0.3? %t (Ложь из-за двоичного представления)\n", sum == f3)

	// Правильное сравнение: с допуском (epsilon)
	epsilon := 1e-10
	diff := math.Abs(sum - f3)
	fmt.Printf("Разница между (0.1+0.2) и 0.3: %.20f\n", diff)
	fmt.Printf("Разница < допуска (1e-10)? %t\n", diff < epsilon)
	fmt.Printf("Сравнение с допуском: %t (Вот это правильный результат)\n", math.Abs(sum-f3) < epsilon)
}