package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	fmt.Println("=== ВСЯ ПРАВДА О FLOAT32 ===")

	// 1. Размер и точность
	var f32 float32 = 123.456
	fmt.Printf("1. Размер и исходное число:\n")
	fmt.Printf("   float32: %f, %T\n", f32, f32)
	fmt.Printf("   Размер: %d байт\n", unsafe.Sizeof(f32))
	// Покажем неточность хранения
	fmt.Printf("   Точное значение: %.15f (Видишь неточность?)\n", f32)
	fmt.Println()

	// 2. Специальные значения: +Inf, -Inf, NaN
	fmt.Println("2. Специальные значения:")
	zero32 := float32(0.0) // Важно: используем float32 ноль
	posInf32 := float32(1.0) / zero32
	negInf32 := float32(-1.0) / zero32
	nan32 := zero32 / zero32

	fmt.Printf("   +Inf: %f\n", posInf32)
	fmt.Printf("   -Inf: %f\n", negInf32)
	fmt.Printf("   NaN: %f\n", nan32)
	fmt.Printf("   NaN == NaN? %t (Всегда false)\n", nan32 == nan32)
	fmt.Println()

	// 3. Проверка специальных значений
	fmt.Println("3. Проверка через math package (работает и для float32):")
	fmt.Printf("   math.IsInf(+Inf, 1): %t\n", math.IsInf(float64(posInf32), 1))
	fmt.Printf("   math.IsInf(-Inf, -1): %t\n", math.IsInf(float64(negInf32), -1))
	fmt.Printf("   math.IsNaN(NaN): %t\n", math.IsNaN(float64(nan32))) // math функции требуют float64
	fmt.Println()

	// 4. Неточность вычислений (еще более выражена, чем у float64!)
	fmt.Println("4. Неточность вычислений (самое интересное):")
	f1_32 := float32(0.1) // Явно указываем тип float32
	f2_32 := float32(0.2)
	f3_32 := float32(0.3)

	fmt.Printf("   f1 (0.1) = %.20f\n", f1_32)
	fmt.Printf("   f2 (0.2) = %.20f\n", f2_32)
	fmt.Printf("   f3 (0.3) = %.20f\n", f3_32)
	sum32 := f1_32 + f2_32
	fmt.Printf("   f1 + f2 (0.1 + 0.2) = %.20f\n", sum32)

	// Прямое сравнение
	fmt.Printf("   0.1 + 0.2 == 0.3? %t (Ошибка еще заметнее!)\n", sum32 == f3_32)

	// Правильное сравнение: с допуском (epsilon). Для float32 эпсилон должен быть больше!
	epsilon32 := float32(1e-6) // Допуск для float32 (1e-6), т.к. точность меньше
	diff32 := float32(math.Abs(float64(sum32 - f3_32))) // Приводим к float64 для math.Abs, потом обратно
	fmt.Printf("   Разница между (0.1+0.2) и 0.3: %.20f\n", diff32)
	fmt.Printf("   Разница < допуска (1e-6)? %t\n", diff32 < epsilon32)
	fmt.Printf("   Сравнение с допуском: %t\n", math.Abs(float64(sum32-f3_32)) < float64(epsilon32))
	fmt.Println()

	// 5. САМОЕ ГЛАВНОЕ: Сравнение точности float32 и float64
	fmt.Println("5. СРАВНЕНИЕ ТОЧНОСТИ float32 vs float64:")
	preciseNum := 123.456789012345
	f32_val := float32(preciseNum)
	f64_val := preciseNum // Это уже float64

	fmt.Printf("   Исходное число:    %.15f\n", preciseNum)
	fmt.Printf("   Значение в float32: %.15f\n", f32_val)
	fmt.Printf("   Значение в float64: %.15f\n", f64_val)
	fmt.Printf("   Потеря точности в float32: %.15f\n", preciseNum - float64(f32_val))
	fmt.Printf("   Потеря точности в float64: %.15f (идеально)\n", preciseNum - f64_val)
	fmt.Println()

	fmt.Println("ВЫВОД: float32 жертвует точностью ради экономии памяти (4 байта вместо 8).")
	fmt.Println("        Для высокой точности ВСЕГДА используй float64.")
}
