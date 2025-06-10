[111. Minimum Depth of Binary Tree](https://leetcode.com/problems/minimum-depth-of-binary-tree/)

```go
package main

type TreeNode struct {
 Val int
 Left *TreeNode
 Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1<<31 - 1

	traverse(root, 1, &depth)

	return depth
}

func traverse(root *TreeNode, level int, minDepth *int) {
	if root == nil {
		return
	}

	if isLeaf(root) && level < *minDepth {
		*minDepth = level
	}

	if level >= *minDepth {
		return
	}

	traverse(root.Left, level+1, minDepth)
	traverse(root.Right, level+1, minDepth)
}

func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}

```

***Оценка по времени:*** `O(n)`, где `n` - количество элементов `TreeNode`

*Объяснения:* делаем 1 проход по всем элементам, который образует `n` итераций

***Оценка по памяти:*** `O(h)`, где `h` - высота дерева

*Объяснения:* в лучшем случае это может быть высота дерева. В худшем пройдем по всем элементам

**Описание решения**

Устанавливаем значение минимальной глубины как максимальное целое. Делаем preorder обход по дереву с передачей 
минимальной глубины и уровня. Если текущая нода это лист и текущий уровень меньше минимальной глубины, то устанавливаем 
минимальную глубину как текущий уровень. Если текущий уровень больше минимальной глубины, то возврат, т.к. нет смысла 
идти далее вниз 

