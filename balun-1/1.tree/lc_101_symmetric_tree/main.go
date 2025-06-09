package main

import "fmt"

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

	fmt.Println(isSymmetric(tree))
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
