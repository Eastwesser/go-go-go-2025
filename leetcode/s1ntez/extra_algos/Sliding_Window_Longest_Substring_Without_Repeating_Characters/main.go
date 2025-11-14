func lengthOfLongestSubstring(s string) int {
    left := 0
    maxLength := 0
    seen := make(map[byte]int)
    
    for right := 0; right < len(s); right++ {
        if idx, exists := seen[s[right]]; exists && idx >= left {
            left = idx + 1 // Сдвигаем окно
        }
        seen[s[right]] = right
        if right-left+1 > maxLength {
            maxLength = right - left + 1
        }
    }
    return maxLength
}