[112. Path Sum](https://leetcode.com/problems/path-sum/)

```go
package main

type TreeNode struct {
 Val int
 Left *TreeNode
 Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return checkHasSum(root, 0, targetSum)
}

func checkHasSum(root *TreeNode, prefixSum int, targetSum int) bool {
	if root == nil {
		return false
	}

	//if root.Left == nil && root.Right == nil && prefixSum+root.Val == targetSum {
	//	return true
	//}

	if isLeaf(root) && prefixSum+root.Val == targetSum {
		return true
	}

	leftHasSum := checkHasSum(root.Left, prefixSum+root.Val, targetSum)
	rightHasSum := checkHasSum(root.Right, prefixSum+root.Val, targetSum)

	return leftHasSum || rightHasSum
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

Делаем preorder обход по дереву с накоплением префиксной суммы. Если корень равен nil, то сразу возврат false. 
Проверяем, что это лист (нет слева и справа нод) и префиксная сумма + значение ноды равно целевой сумме. Далее, 
рекурсивно идем влево и потом вправо

