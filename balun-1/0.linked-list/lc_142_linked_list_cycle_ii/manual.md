[142. Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/)

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
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			slow = head

			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}

			return slow
		}
	}

	return nil
}
```

***Оценка по времени:*** `O(n)`, где `n` - количество элементов `ListNode`

*Объяснения:* делаем 2 прохода по всем элементам, который образует `2n` итераций. Константы при `O` можно опускать.

***Оценка по памяти:*** `O(1)`

*Объяснения:* количество элементов неизменнно

**Описание решения**

Необходимо применить алгоритм Флойда ("черепаха и заяц"). Создаем 2 указателя: slow и fast. Оба устанавливаем в head.
В цикле пока fast и fast.Next не равны nil перемещаем указатели: slow на 1 ноду, fast - на 2. Если внутри цикла нет, то
вернется nil. Если есть, то slow встретится с fast рано или поздно. Тогда устанавливаем slow в head и в цикле идем смещаясь на 1 ноду, 
пока они не встретятся. Возвращаем slow (или fast)

[Алгоритм Флойда](https://gabhisekdev.medium.com/why-does-floyds-cycle-detection-algorithm-work-59f61984dc3e)