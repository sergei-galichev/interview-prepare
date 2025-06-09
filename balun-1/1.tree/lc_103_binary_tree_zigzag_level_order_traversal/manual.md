[103. Binary Tree Zigzag Level Order Traversal](https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/)

```go
package main

type TreeNode struct {
 Val int
 Left *TreeNode
 Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)

	preorder(root, 0, &result)

	for i := 0; i < len(result); i++ {
		if i%2 == 1 {
			for j := 0; j < len(result[i])/2; j++ {
				result[i][j], result[i][len(result[i])-j-1] = result[i][len(result[i])-j-1], result[i][j]
			}
		}
	}

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

***Оценка по памяти:*** `O(n)`, где `n` - количество элементов `TreeNode`

*Объяснения:* количество элементов неизменнно

**Описание решения**

Делаем обход как при level order, т.е. формируем двумерный слайс. Далее, каждый нечетный слайс внутри него разворачиваем