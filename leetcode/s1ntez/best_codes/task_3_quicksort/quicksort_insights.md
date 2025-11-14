# ✅ Generic-подход:
```go
func QuickSort(slice []User, less func(a, b User) bool)
```
Можно сортировать по любому полю и любому порядку!

# ✅ Гибкость сортировки:
    По ID: a.ID < b.ID
    По Age: a.Age > b.Age (descending!)
    По Name: a.Name < b.Name
    Или даже комбинации!

# ✅ In-place сортировка:
Не создаёт новый слайс - эффективно по памяти!
