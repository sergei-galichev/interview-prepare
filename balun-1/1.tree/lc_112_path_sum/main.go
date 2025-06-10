package main

import (
	"fmt"
)

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 5,
			},
			//Right: &TreeNode{
			//	Val: 4,
			//},
		},
	}

	fmt.Println(hasPathSum(tree, 7))
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
