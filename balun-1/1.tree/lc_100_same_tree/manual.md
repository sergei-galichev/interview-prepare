[100. Same Tree](https://leetcode.com/problems/same-tree/)

```go
package main

type TreeNode struct {
 Val int
 Left *TreeNode
 Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

```

***Оценка по времени:*** `O(n)`, где `n` - количество элементов `TreeNode`

*Объяснения:* делаем 1 проход по всем элементам, который образует `n` итераций

***Оценка по памяти:*** `O(h)`, где `h` - высота дерева

*Объяснения:* в лучшем случае это может быть высота дерева. В худшем пройдем по всем элементам

**Описание решения**

Делаем рекурсивный обход 2 деревьев. Проверяем, если один из указателей равен nil, то проверяем 2-й на это равенство. 
Далее, если значения узлов не равны, то возврат false

