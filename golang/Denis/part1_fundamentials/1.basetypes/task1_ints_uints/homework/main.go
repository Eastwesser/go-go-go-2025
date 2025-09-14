package main

// Silver Bullet: XOR или формула суммы арифметической прогрессии
func FindMissingNumber(nums []int) int {
    n := len(nums)
    
	total := n * (n + 1) / 2
    sum := 0
    
	for _, num := range nums {
        sum += num
    }
    
	return total - sum
}

// Chaos Version: С дубликатами и отрицательными числами
func FindMissingNumberChaos(nums []int) int {
    seen := make(map[int]bool)
    for _, num := range nums {
        if num < 0 || seen[num] {
            return -1 // Аномалия!
        }
        seen[num] = true
    }

	res := int(seen[0])
    return res
}

func main() {

}