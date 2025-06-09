[101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree/)

```go
package main

type TreeNode struct {
 Val int
 Left *TreeNode
 Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return checkSymmetry(root.Left, root.Right)
}

func checkSymmetry(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == nil && right == nil
	}

	if left.Val != right.Val {
		return false
	}

	return checkSymmetry(left.Left, right.Right) && checkSymmetry(left.Right, right.Left)
}

```

***Оценка по времени:*** `O(n)`, где `n` - количество элементов `TreeNode`

*Объяснения:* делаем 1 проход по всем элементам, который образует `n` итераций

***Оценка по памяти:*** `O(h)`, где `h` - высота дерева

*Объяснения:* в лучшем случае это может быть высота дерева. В худшем пройдем по всем элементам

**Описание решения**

Если корень равен nil, то возврат false. Далее передаем 2 указателя: левый и правый. Рекурсивно идем по левому и
правому деревьям. Если указатель левый либо правый равен nil, то проверям равны ли оба nil и возвращаем true, если это
так. Далее проверяем значения. Если не равны, то возврат false

