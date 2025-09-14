// Silver Bullet: Суммирование с Kahan компенсацией
func CalculateAverage(nums []float64) float64 {
    var sum, compensation float64
    for _, num := range nums {
        y := num - compensation
        t := sum + y
        compensation = (t - sum) - y
        sum = t
    }
    return sum / float64(len(nums))
}

// Chaos Version: С бесконечностями и NaN
func CalculateAverageChaos(nums []float64) float64 {
    var validNums []float64
    for _, num := range nums {
        if !math.IsNaN(num) && !math.IsInf(num, 0) {
            validNums = append(validNums, num)
        }
    }
    if len(validNums) == 0 {
        return math.NaN()
    }
    // Вычисление среднего...
}