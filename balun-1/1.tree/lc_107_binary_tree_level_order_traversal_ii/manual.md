[107. Binary Tree Level Order Traversal II](https://leetcode.com/problems/binary-tree-level-order-traversal-ii/)

```go
package main

import (
	"slices"
)

type TreeNode struct {
 Val int
 Left *TreeNode
 Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)

	postorder(root, 0, &result)

	return result
}

func postorder(root *TreeNode, level int, result *[][]int) {
	if root == nil {
		return
	}

	if level == len(*result) {
		*result = slices.Insert(*result, 0, make([]int, 0))
	}

	postorder(root.Left, level+1, result)
	postorder(root.Right, level+1, result)

	(*result)[len(*result)-level-1] = append((*result)[len(*result)-level-1], root.Val)
}

```

***Оценка по времени:*** `O(n)`, где `n` - количество элементов `TreeNode`

*Объяснения:* делаем 1 проход по всем элементам, который образует `n` итераций

***Оценка по памяти:*** `O(n)`, где `n` - количество элементов `TreeNode`.

*Объяснения:* количество элементов неизменнно

**Описание решения**

Делаем обход postorder. В результирующий слайс добавляем новый слайс элементов для текущего уровня в начало 
(`slices.Insert`), если он ранее не добавлялся. Идем влево и вправо рекурсивно. После добавляем `root.Val` в слайс с 
индексом `len(*result)-level-1`
