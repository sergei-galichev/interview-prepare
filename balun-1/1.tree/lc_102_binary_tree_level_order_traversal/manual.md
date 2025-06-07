[102. Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/)

```go
package main

type TreeNode struct {
 Val int
 Left *TreeNode
 Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)

	preorder(root, 0, &result)

	return result
}

func preorder(root *TreeNode, level int, result *[][]int) {
	if root == nil {
		return
	}

	if level == len(*result) {
		*result = append(*result, make([]int, 0))
	}

	(*result)[level] = append((*result)[level], root.Val)

	preorder(root.Left, level+1, result)
	preorder(root.Right, level+1, result)
}

```

***Оценка по времени:*** `O(n)`, где `n` - количество элементов `TreeNode`

*Объяснения:* делаем 1 проход по всем элементам, который образует `n` итераций

***Оценка по памяти:*** `O(h)`, где `h` - высота дерева

*Объяснения:* количество элементов неизменнно

**Описание решения**

Пишем вспомогательную функцию. В нее передаем корень, уровень и двумерный слайс. Если корень nil, то выход. Если длина
результирующего слайса равна уровню, то добавляем еще слайс в результирующий двумерный слайс. Добавляем значения узла в
результирующий слайс согласно уровню. Потом идем влево рекурсивно. После идем вправо рекурсивно