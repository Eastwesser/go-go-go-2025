// Silver Bullet: Подсчёт символов через массив [26]int
func IsAnagram(s1, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    count := [26]int{}
    for i := 0; i < len(s1); i++ {
        count[s1[i]-'a']++
        count[s2[i]-'a']--
    }
    for _, c := range count {
        if c != 0 {
            return false
        }
    }
    return true
}

// Chaos Version: Unicode и специальные символы
func IsAnagramChaos(s1, s2 string) bool {
    return normalizeString(s1) == normalizeString(s2)
}

func normalizeString(s string) string {
    var result []rune
    for _, r := range strings.ToLower(s) {
        if unicode.IsLetter(r) {
            result = append(result, r)
        }
    }
    sort.Slice(result, func(i, j int) bool {
        return result[i] < result[j]
    })
    return string(result)
}