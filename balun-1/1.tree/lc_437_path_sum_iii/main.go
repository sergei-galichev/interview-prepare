package main

import (
	"fmt"
)

func main() {
	//tree := &TreeNode{
	//	Val: 10,
	//	Left: &TreeNode{
	//		Val: 5,
	//		Left: &TreeNode{
	//			Val: 3,
	//			Left: &TreeNode{
	//				Val: 3,
	//			},
	//			Right: &TreeNode{
	//				Val: -2,
	//			},
	//		},
	//		Right: &TreeNode{
	//			Val: 2,
	//			Right: &TreeNode{
	//				Val: 1,
	//			},
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val: -3,
	//		Right: &TreeNode{
	//			Val: 11,
	//		},
	//	},
	//}

	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}

	fmt.Println(pathSum(tree, 22))
	fmt.Println(pathSumOptimized(tree, 22))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
time: O(n), n is node count
mem: O(h) h is height (h may be n)
*/

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
