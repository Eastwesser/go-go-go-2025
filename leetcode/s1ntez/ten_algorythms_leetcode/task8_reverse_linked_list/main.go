package main

import "fmt"

type ListNode8 struct {
	Val int
	Next *ListNode8
}

func ReverseLinkedList(head *ListNode8) *ListNode8 {
	var prev *ListNode8
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

// Функция для красивого вывода списка
func printList(head *ListNode8) {
	current := head
	for current != nil {
		fmt.Printf("%d", current.Val)
		if current.Next != nil {
			fmt.Print(" → ")
		}
		current = current.Next
	}
	fmt.Println()
}

// Функция для создания списка из слайса (удобно для тестов)
func createList(values []int) *ListNode8 {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode8{Val: values[0]}
	current := head
	for i := 1; i < len(values); i++ {
		current.Next = &ListNode8{Val: values[i]}
		current = current.Next
	}
	return head
}

func main() {
	// Создаем список: 1 → 3 → 5
	list := createList([]int{1, 3, 5})
	
	fmt.Print("Исходный список: ")
	printList(list) // 1 → 3 → 5

	// Разворачиваем
	reversed := ReverseLinkedList(list)
	
	fmt.Print("Развернутый список: ")
	printList(reversed) // 5 → 3 → 1
	
	// Дополнительная проверка
	fmt.Print("Еще один тест: ")
	list2 := createList([]int{1, 2, 3, 4, 5})
	printList(list2) // 1 → 2 → 3 → 4 → 5
	printList(ReverseLinkedList(list2)) // 5 → 4 → 3 → 2 → 1
}
