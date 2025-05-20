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

func reorderList(head *ListNode) {
    tmp := preMiddle(head)

    ptr2 := reverseList(tmp.Next)

    tmp.Next = nil

    ptr1 := head

    for ptr2 != nil {
        ptr1Next := ptr1.Next
        ptr1.Next = ptr2
        ptr1 = ptr2
        ptr2 = ptr1Next
    }
}

func preMiddle(head *ListNode) *ListNode {
    slow, fast := head, head

    for fast != nil && fast.Next != nil && fast.Next.Next != nil {
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

Найти pre-middle: для списка с четным количеством это будет элемент перед middle, с четным - это middle элемент.
Далее передаем preMiddle.Next для разворота списка.  Он вернет ptr2. Указатель ptr1 ставим на head. Пока ptr2 не nil,
меняем ptr1, чтобы указывал на элемент ptr2, а ptr2 - на элемент ptr1.Next (до изменения ptr1).
