[199. Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view/)

```go
package main

type TreeNode struct {
 Val int
 Left *TreeNode
 Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)

	levelOrder(root, 0, &result)

	return result
}

func levelOrder(root *TreeNode, level int, result *[]int) {
	if root == nil {
		return
	}

	if level == len(*result) {
		*result = append(*result, 0)
	}

	(*result)[level] = root.Val
	levelOrder(root.Left, level+1, result)
	levelOrder(root.Right, level+1, result)
}

```

***Оценка по времени:*** `O(n)`, где `n` - количество элементов `TreeNode`

*Объяснения:* делаем 1 проход по всем элементам, который образует `n` итераций

***Оценка по памяти:*** `O(h)`, где `h` - высота дерева

*Объяснения:* количество элементов неизменнно

**Описание решения**

Используем обход 'level order' за тем лишь исключением, что результат в одномерном массиве. Значения на одном уровне
перезаписываются значением последнего посещенного узла

