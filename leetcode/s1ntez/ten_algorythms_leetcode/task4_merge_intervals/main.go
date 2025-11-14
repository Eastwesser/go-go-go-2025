package main

import (
	"fmt"
	"sort"
)

func MergeIntervals(intervals [][]int) [][]int {
    // Базовый случай: если 0 или 1 интервал - нечего мерджить
    if len(intervals) <= 1 {
        return intervals
    }

    // Сортируем интервалы по начальной точке (intervals[i][0])
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    // Начинаем с первого интервала
    res := [][]int{intervals[0]}

    // Проходим по всем остальным интервалам
    for i := 1; i < len(intervals); i++ {
        last := res[len(res)-1]  // Последний интервал в результате
        current := intervals[i]  // Текущий интервал

        // Проверяем пересечение: начало текущего <= конец последнего
        if current[0] <= last[1] {
            // Если пересекаются - расширяем последний интервал
            if current[1] > last[1] {
                last[1] = current[1]  // Обновляем конец интервала
            }
        } else {
            // Не пересекаются - добавляем новый интервал
            res = append(res, current)
        }
    }

    return res
}

func main() {
    tests := []struct {
        name string
        input [][]int
        expected [][]int
    }{
        {
            name: "Basic merge",
            input: [][]int{{1,3}, {2,6}, {8,10}, {15,18}},
            expected: [][]int{{1,6}, {8,10}, {15,18}},
        },
        {
            name: "Complete overlap", 
            input: [][]int{{1,4}, {2,3}},
            expected: [][]int{{1,4}},
        },
        {
            name: "No overlaps",
            input: [][]int{{1,2}, {3,4}, {5,6}},
            expected: [][]int{{1,2}, {3,4}, {5,6}},
        },
    }

    for _, test := range tests {
        result := MergeIntervals(test.input)
        fmt.Printf("Test: %s\n", test.name)
        fmt.Printf("Input: %v\n", test.input)
        fmt.Printf("Expected: %v\n", test.expected) 
        fmt.Printf("Result: %v\n", result)
        fmt.Println("---")
    }
}
