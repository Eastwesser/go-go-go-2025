package main

import "testing"

func BenchmarkMergeTwoLists(b *testing.B) {
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        // Создаём свежие списки на каждой итерации
        l1 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}
        l2 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6}}}
        MergeTwoLists(l1, l2)
        
        // Тест с пустым списком
        l3 := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
        var l4 *ListNode
        MergeTwoLists(l3, l4)
    }
}

// go test -bench=. -benchmempackage main
