// Silver Bullet: Рекурсивный QuickSort с выбором опоры
func QuickSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }
    pivot := arr[len(arr)/2]
    var left, right, equal []int
    for _, num := range arr {
        switch {
        case num < pivot:
            left = append(left, num)
        case num == pivot:
            equal = append(equal, num)
        default:
            right = append(right, num)
        }
    }
    return append(append(QuickSort(left), equal...), QuickSort(right)...)
}

// Chaos Version: С защитой от худшего случая и stack overflow
func QuickSortChaos(arr []int, maxDepth int) []int {
    if maxDepth == 0 {
        return heapSort(arr) // Переключаемся на heapsort
    }
    // Рекурсивная логика...
}