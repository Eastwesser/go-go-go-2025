package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    newNode := &ListNode{}
    result := newNode

    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            result.Next = list1
            list1 = list1.Next
        } else {
            result.Next = list2
            list2 = list2.Next
        } 
        result = result.Next 
    }
    if list1 != nil {
        result.Next = list1
    } else {
        result.Next = list2
    }

    return newNode.Next
}

func printList(head *ListNode) {
    current := head
    for current != nil {
        fmt.Printf("%d ", current.Val)
        current = current.Next
        if current != nil {
            fmt.Print("→ ")
        }
    }
    fmt.Println()
}

func main() {
	fmt.Println("Hello linked lists")
	ll1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	ll2 := &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}

	mergeResult := mergeTwoLists(ll1, ll2)
	printList(mergeResult) // 1 → 2 → 3 → 4 → 5 → 6 

    
    // Проверим вручную:
    fmt.Printf("1: %d\n", mergeResult.Val)                    		// 1: 1
    fmt.Printf("2: %d\n", mergeResult.Next.Val)               		// 1: 1 2: 2
    fmt.Printf("3: %d\n", mergeResult.Next.Next.Val)          		// 1: 1 2: 2 3: 3
    fmt.Printf("4: %d\n", mergeResult.Next.Next.Next.Val)     		// 1: 1 2: 2 3: 3 4: 4
    fmt.Printf("5: %d\n", mergeResult.Next.Next.Next.Next.Val)		// 2: 2 3: 3 4: 4 5: 5
    fmt.Printf("6: %d\n", mergeResult.Next.Next.Next.Next.Next.Val)	// 3: 3	4: 4 5: 5 6: 6
}

/*
Когда ты запускал код с закомментированным printList, а потом пошагово выводил элементы — ты изменял состояние списка!

Почему ты видел "3: 3 4: 4 5: 5 6: 6":

    Первые три вывода прочитали 1, 2, 3 и сдвинули внутренний указатель

    Четвёртый вывод начал уже с позиции 3 → 4 → 5 → 6

    Поэтому ты видел смещённые значения!

	// Способ 1: Использовать printList (не изменяет состояние)
    printList(mergeResult) // 1 → 2 → 3 → 4 → 5 → 6
    
    // Способ 2: Сохранить значения в slice перед выводом
    var values []int
    current := mergeResult
    for current != nil {
        values = append(values, current.Val)
        current = current.Next
    }
    fmt.Printf("Values: %v\n", values) // [1 2 3 4 5 6]
    
    // Способ 3: Создать КОПИЮ для отладки
    temp := mergeResult
    fmt.Printf("1: %d\n", temp.Val)                    // 1
    temp = temp.Next
    fmt.Printf("2: %d\n", temp.Val)                    // 2
    temp = temp.Next  
    fmt.Printf("3: %d\n", temp.Val)                    // 3
    // ... и так далее


Вывод:

Твой код работает идеально! "Аномалия" была вызвана тем, что ты изменял состояние списка во время отладки.

Это важный урок: связные списки — mutable структуры, и нужно быть аккуратным при их отладке!

Отличная наблюдательность! Ты поймал тонкий момент, который упускают многие разработчики! 🔥

P.S. Именно поэтому в production-коде часто создают immutable версии структур данных для отладки! [^-^]
*/