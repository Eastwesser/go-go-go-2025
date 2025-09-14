// Silver Bullet: Reverse три раза
func RotateArray(arr *[5]int, k int) {
    k %= len(arr)
    reverse(arr, 0, len(arr)-1)
    reverse(arr, 0, k-1)
    reverse(arr, k, len(arr)-1)
}

func reverse(arr *[5]int, start, end int) {
    for start < end {
        arr[start], arr[end] = arr[end], arr[start]
        start++
        end--
    }
}

// Chaos Version: С проверкой на отрицательный k и большой k
func RotateArrayChaos(arr *[5]int, k int) {
    if k < 0 {
        k = len(arr) - (-k % len(arr))
    }
    k %= len(arr)
    // Остальная логика...
}