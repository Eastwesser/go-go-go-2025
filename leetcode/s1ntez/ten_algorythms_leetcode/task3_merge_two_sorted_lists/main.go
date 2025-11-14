package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func MergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1 
	} else {
		current.Next = l2
	}

	return dummy.Next
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d → ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	// Создаем два отсортированных списка
	// list1: 1 → 3 → 5
	list1 := &ListNode{
		Val: 1, 
		Next: &ListNode{
			Val: 3, 
			Next: &ListNode{
				Val: 5,
			},
		},
	}

	// list2: 2 → 4 → 6  
	list2 := &ListNode{
		Val: 2, 
		Next: &ListNode{
			Val: 4, 
			Next: &ListNode{
				Val: 6,
			},
		},
	}

	fmt.Print("List 1: ")
	printList(list1)
	fmt.Print("List 2: ")
	printList(list2)

	// Мерджим
	merged := MergeTwoLists(list1, list2)
	fmt.Print("Merged: ")
	printList(merged) // Должно быть: 1 → 2 → 3 → 4 → 5 → 6 → nil
}
