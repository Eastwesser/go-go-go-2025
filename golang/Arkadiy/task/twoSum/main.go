package main

import "fmt"

func twoSum(slise []int, target int) []int {

	for i := 0; i < len(slise); i++ {
		for j := 0; j < len(slise); j++ {
			if target == slise[i]+slise[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	array := []int{5, 4, 6}
	result := twoSum(array, 11)
	fmt.Println(result)
}
