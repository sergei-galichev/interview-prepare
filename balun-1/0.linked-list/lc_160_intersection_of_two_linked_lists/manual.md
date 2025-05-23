[160. Intersection of Two Linked Lists](https://leetcode.com/problems/intersection-of-two-linked-lists/)

```go
package main

type ListNode struct {
    Val  int
    Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	lengthA := listLength(headA)
	lengthB := listLength(headB)

	ptrA, ptrB := headA, headB

	if lengthA > lengthB {
		for i := 0; i < lengthA-lengthB; i++ {
			ptrA = ptrA.Next
		}
	} else if lengthA < lengthB {
		for i := 0; i < lengthB-lengthA; i++ {
			ptrB = ptrB.Next
		}
	}

	for ptrA != nil && ptrB != nil {
		if ptrA == ptrB && ptrA.Val == ptrB.Val {
			return ptrA
		}

		ptrA = ptrA.Next
		ptrB = ptrB.Next
	}

	return nil
}

func listLength(head *ListNode) int {
	length := 0

	curr := head

	for curr != nil {
		curr = curr.Next

		length++
	}

	return length
}
```

***Оценка по времени:*** `O(n+m)`, где `n` - количество элементов `ListNode` в headA, `m` - количество элементов 
`ListNode` в headB. 

*Объяснения:* делаем 1 проход по всем элементам каждого списка, который образует `n` и `m` итераций соответственно.
Выравнивание указателей в худшем случае займет `O(|n-m|)` шагов. После выравнивания обход идет до конца списка
`O(min(n,m))` 

***Оценка по памяти:*** `O(1)`

*Объяснения:* количество элементов неизменнно

**Описание решения**

Проверка, что указатели на начала списков не равны nil. Сначала нужно выровнять указатели. Для этого вычисляем длины списков и в зависимости какой список длиннее, тот указатель
и смещаем. Потом идем по обоим спискам и сравниваем ссылки на ноды и значения. В случае совпадения возврат указателя на 
ноду. В противном случае вернется nil.

