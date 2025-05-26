[206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)

```go
package main

type ListNode struct {
    Val  int
    Next *ListNode
}

/*
*
  - Definition for singly-linked list.
  - type ListNode struct {
  - Val int
  - Next *ListNode
  - }

time: O(n), n is node count
mem: O(1)
*/

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

```

***Оценка по времени:*** `O(n)`, где `n` - количество элементов `ListNode`

*Объяснения:* делаем 1 проход по всем элементам, который образует `n` итераций

***Оценка по памяти:*** `O(1)`

*Объяснения:* количество элементов неизменнно

**Описание решения**

Создаем два указателя slow и fast. И в цикле идем по всем нодам пока fast и fast.Next не равны nil. Т.к. быстрый
указать fast смещается на 2 ноды, а slow только на 1, то он догонит slow через некоторое время, если есть цикл.

