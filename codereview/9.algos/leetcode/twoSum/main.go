package main

import (
    "fmt"
)

// Нужно найти максимальную сумму чисел из двух массивов
func TwoSum(a, b []int) int {
    // найти максимальную сумму
    maxSumOfFirst := a[0]
    for _, num := range a {
       if num > maxSumOfFirst {
          maxSumOfFirst = num
       }
    }

    maxSumOfSecond := b[0]
    for _, num := range b {
       if num > maxSumOfSecond {
          maxSumOfSecond = num
       }
    }

    return maxSumOfFirst + maxSumOfSecond
}

func main() {
    numArray1 := []int{2, 5, 9}
    numArray2 := []int{3, 6, 10}
    fmt.Println(TwoSum(numArray1, numArray2))
}
