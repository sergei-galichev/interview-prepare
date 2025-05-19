[234. Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/)

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
func isPalindrome(head *ListNode) bool {
	firstHalf := middleNode(head)

	secondHalf := reverseList(firstHalf)

	ptr1 := head
	ptr2 := secondHalf

	for ptr2 != nil {
		if ptr1.Val != ptr2.Val {
			return false
		}

		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	return true
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	curr := head

	for curr != nil {
		tmp := curr

		curr = curr.Next

		tmp.Next = prev

		prev = tmp
	}

	return prev
}
```

***Оценка по времени:*** `O(n)`, где `n` - количество элементов `ListNode`

*Объяснения:* делаем 1 проход по всем элементам, который образует `n` итераций

***Оценка по памяти:*** `O(1)`

*Объяснения:* количество элементов неизменнно

**Описание решения**

Сначала находим середину через slow и fast указатели. Потом делаем реверс второй половины и начинаем смещать указатели 
prt1 и ptr2 пока ptr2 не равен nil. Сравниваем значения Val обоих указателей. Если не равны, то false