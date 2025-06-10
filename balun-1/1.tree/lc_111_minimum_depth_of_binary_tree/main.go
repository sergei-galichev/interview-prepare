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
				Left: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	fmt.Println(minDepth(tree))
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
