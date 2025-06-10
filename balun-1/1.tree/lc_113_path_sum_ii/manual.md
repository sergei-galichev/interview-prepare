[113. Path Sum II](https://leetcode.com/problems/path-sum-ii/)

```go
package main

type TreeNode struct {
 Val int
 Left *TreeNode
 Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := make([][]int, 0)
	prefixSlice := make([]int, 0)

	checkHasSum(root, 0, targetSum, &prefixSlice, &result)

	return result
}

func checkHasSum(root *TreeNode, prefixSum, targetSum int, prefixSlice *[]int, result *[][]int) {
	if root == nil {
		return
	}

	*prefixSlice = append(*prefixSlice, root.Val)
	defer func() {
		*prefixSlice = (*prefixSlice)[:len(*prefixSlice)-1]
	}()

	if isLeaf(root) && prefixSum+root.Val == targetSum {
		*result = append(*result, append([]int{}, *prefixSlice...))

		return
	}

	checkHasSum(root.Left, prefixSum+root.Val, targetSum, prefixSlice, result)
	checkHasSum(root.Right, prefixSum+root.Val, targetSum, prefixSlice, result)
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

Делаем тот же preorder обход по дереву с накоплением префиксной суммы. Т.к. нужно найти все пути соответствующие условию,
то передаем рекурсивно 2 слайса: результирующий и буферный. При движении вниз по дереву влево или вправо в буферный 
слайс добавляем элементы. При возврате обрезаем буферный слайс на 1 элемент с конца 

