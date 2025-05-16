[19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)

```go
package main

type ListNode struct {
    Val  int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummyNode := &ListNode{0, head}
    
    length := 0
    
    curr := dummyNode
    
    for curr != nil {
    curr = curr.Next
    
    length++
    }
    
    curr = dummyNode
    
    for i := 0; i < length-n-1; i++ {
    curr = curr.Next
    }
    
    curr.Next = curr.Next.Next
    
    return dummyNode.Next
}
```

***Оценка по времени:*** `O(n)`, где `n` - количество элементов `ListNode`

*Объяснения:* делаем 1 проход по всем элементам, который образует `n` итераций

***Оценка по памяти:*** `O(1)`

*Объяснения:* количество элементов неизменнно

**Описание решения**

Используем паттерн DummyNode