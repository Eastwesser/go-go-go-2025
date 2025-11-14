package main

import "testing"

func BenchmarkReverseLinkedList(b *testing.B) {
	createList := func(size int) *ListNode8 {
		head := &ListNode8{Val: 1}
		current := head
		for i := 2; i <= size; i++ {
			current.Next = &ListNode8{Val: i}
			current = current.Next
		}
		return head
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list := createList(111)
		ReverseLinkedList(list)
	}
}
/*
BenchmarkReverseLinkedList-4       46052             55497 ns/op            1776 B/op        111 allocs/op
*/
