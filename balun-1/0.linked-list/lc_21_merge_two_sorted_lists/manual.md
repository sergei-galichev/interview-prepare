[21. Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)

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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyNode := &ListNode{0, nil}

	stab := dummyNode

	for list1 != nil || list2 != nil {
		if getNodeValue(list1) < getNodeValue(list2) {
			dummyNode.Next = list1
			list1 = list1.Next
		} else {
			dummyNode.Next = list2
			list2 = list2.Next
		}

		dummyNode = dummyNode.Next
	}

	return stab.Next
}

func getNodeValue(node *ListNode) float64 {
	if node == nil {
		return math.Inf(1)
	}

	return float64(node.Val)
}

```

***Оценка по времени:*** `O(n)`, где `n` - количество элементов `ListNode`

*Объяснения:* делаем 1 проход по всем элементам, который образует `n` итераций

***Оценка по памяти:*** `O(1)`

*Объяснения:* количество элементов неизменнно

**Описание решения**

Используем DummyNode. Добавляем функцию getNodeValue, которая возвращает значение Val, если node не равна nil. В
противном случаем вернется бесконечность. Создаем DummyNode и делаем stab, указывающий на нее. Далее по циклу сравниваем
значения ListNode и устанавливаем указатели на Next, пока list1 или list2 не равен nil.