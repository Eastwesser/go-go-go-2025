// Silver Bullet: Итерация по меньшему мапу
func MapIntersection(m1, m2 map[string]int) map[string]int {
    result := make(map[string]int)
    // Выбираем меньший мап для итерации
    if len(m1) > len(m2) {
        m1, m2 = m2, m1
    }
    for k, v1 := range m1 {
        if v2, exists := m2[k]; exists {
            result[k] = min(v1, v2)
        }
    }
    return result
}

// Chaos Version: С обработкой nil мапов и разных типов
func MapIntersectionChaos(m1, m2 interface{}) interface{} {
    // Рефлексия для обработки разных типов...
}