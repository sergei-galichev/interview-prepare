[94. Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/)

```go
package main

type TreeNode struct {
 Val int
 Left *TreeNode
 Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	traverse(root, &result)

	return result
}

func traverse(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	traverse(root.Left, result)
	*result = append(*result, root.Val)
	traverse(root.Right, result)
}

```

***Оценка по времени:*** `O(n)`, где `n` - количество элементов `TreeNode`

*Объяснения:* делаем 1 проход по всем элементам, который образует `n` итераций

***Оценка по памяти:*** `O(h)`, где `h` - высота дерева

*Объяснения:* количество элементов неизменнно

**Описание решения**

Пишем дополнительную функцию. Если переданный корень nil, то возврат. Сначала идем рекурсивно влево. Потом добавляем
значение в результирующий массив. Потом рекурсивно вправо

