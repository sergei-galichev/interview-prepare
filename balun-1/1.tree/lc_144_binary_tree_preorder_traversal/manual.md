[144. Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/)

```go
package main

type TreeNode struct {
 Val int
 Left *TreeNode
 Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	traverse(root, &result)

	return result
}

func traverse(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)

	traverse(root.Left, result)
	traverse(root.Right, result)
}

```

***Оценка по времени:*** `O(n)`, где `n` - количество элементов `TreeNode`

*Объяснения:* делаем 1 проход по всем элементам, который образует `n` итераций

***Оценка по памяти:*** `O(h)`, где `h` - высота дерева

*Объяснения:* количество элементов неизменнно

**Описание решения**

Пишем дополнительную функцию. Если переданный корень nil, то возврат. Добавляем значение в результирующий массив. Потом
идем влево рекурсивно - передаем узел слева. Потом рекурсивно вправо

