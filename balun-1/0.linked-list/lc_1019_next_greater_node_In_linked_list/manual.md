[1019. Next Greater Node In Linked List](https://leetcode.com/problems/next-greater-node-in-linked-list/)

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
mem: O(n)
*/
func nextLargerNodes(head *ListNode) []int {
	nums := make([]int, 0)
	curr := head

	for curr != nil {
		nums = append(nums, curr.Val)
		curr = curr.Next
	}

	stack := make([]int, 0)
	result := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			result[i] = stack[len(stack)-1]
		}

		stack = append(stack, nums[i])
	}

	return result
}
```

***Оценка по времени:*** `O(n)`, где `n` - количество элементов `ListNode`

*Объяснения:* делаем 2 прохода по всем элементам, который образует `2n` итераций. Константы при `O` можно опускать.

***Оценка по памяти:*** `O(n)`

*Объяснения:* количество элементов неизменнно

**Описание решения**

Сначала создаем слайс из значений нод и заполняем его значениями, пройдя по всем нодам. Создаем также слайс нулевой
длины (стек) и результирующий, заполненный 0. Идем по слайсу значений в обратном порядке в цикле. Внутри цикла проверяем
что стек не пуст и текущий элемент больше либо равен элементу на вершине стека. Тогда элемент удаляется из стека. Далее,
если стек не пуст, то элемент на вершине стека устанавливаем в результирующий слайс по текущему индексу. В стек
добавляем текущий элемент из слайса значений