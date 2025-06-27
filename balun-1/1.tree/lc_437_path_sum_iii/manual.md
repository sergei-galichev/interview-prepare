[437. Path Sum III](https://leetcode.com/problems/path-sum-iii/)

```go
package main

type TreeNode struct {
 Val int
 Left *TreeNode
 Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	return checkAndCount(root, targetSum) + checkAndCount(root.Left, targetSum) + checkAndCount(root.Right, targetSum)
}

func checkAndCount(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	countPaths := 0

	if root.Val == targetSum {
		countPaths++
	}

	countPaths += checkAndCount(root.Left, targetSum-root.Val)
	countPaths += checkAndCount(root.Right, targetSum-root.Val)

	return countPaths
}

func pathSumOptimized(root *TreeNode, targetSum int) int {
	prefixSums := map[int]int{0: 1}

	return checkHasSumOptimized(root, 0, targetSum, prefixSums)
}

func checkHasSumOptimized(root *TreeNode, currentSum, targetSum int, prefixSums map[int]int) int {
	if root == nil {
		return 0
	}

	currentSum += root.Val
	countPaths := 0

	countPaths += prefixSums[currentSum-targetSum]

	prefixSums[currentSum]++

	countPaths += checkHasSumOptimized(root.Left, currentSum, targetSum, prefixSums)
	countPaths += checkHasSumOptimized(root.Right, currentSum, targetSum, prefixSums)

	prefixSums[currentSum]--

	return countPaths
}

```

***Оценка по времени:*** `O(n)`, где `n` - количество элементов `TreeNode`

*Объяснения:* делаем 1 проход по всем элементам, который образует `n` итераций

***Оценка по памяти:*** `O(n)`, где `n` - количество элементов `TreeNode`

*Объяснения:* делаем 1 проход по всем элементам, который образует `n` итераций

**Описание решения**

Решение в лоб
Делаем 3 preorder обхода по дереву с уменьшением `targetSum` на `root.Val`. Сначала начинаем от корня, потом слева и 
наконец справа от корня. На каждом последующем шаге проверяем `root.Val == targetSum`. В конце все полученные результаты
суммируем

Решение с мапой префиксных сумм
Делаем preorder обход по дереву. Перед началом обхода создаем `prefixSums := map[int]int{0: 1}`. Делаем подсчет текущей
суммы. Из мапы префикных сумм пытаемся получить по ключу значение и сложить с счетчиком найденных путей. Ключом является
разность текущей и целевой сумм. В мапе делаем инкремент значения по ключу, который является текущей суммой. И теперь 
рекурсивно идем влево и затем вправо. Результаты складываем с счетчиком найденных путей. Делаем декремент значения 
по ключу, который является текущей суммой, т.к. в вышестоящих операциях данная префиксная сумма не должна учитываться.
И возвращаем счетчик найденных путей.

