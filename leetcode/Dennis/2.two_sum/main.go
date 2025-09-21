package main

import "fmt"

func FindTwoSum(nums []int) (int, error) {
	left, right := 0, len(nums) - 1
	counter := 0
	last := 0
	next := 0

	for left < right {
		for i := 0; i < len(nums); i++ {
			if left == right {
				counter = nums[last] + nums[next]
				return counter, nil
			} else {
				return counter, nil
			}
		}	
	}

	return counter, nil
}

func main() {
	fmt.Println("Two Sum")

	numList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(FindTwoSum(numList))
}
