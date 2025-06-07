package main

import "fmt"

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(postorderTraversal(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
time: O(n), n is node count
mem: O(h) h is height of tree
*/

func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	traverse(root, &result)

	return result
}

func traverse(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	traverse(root.Left, result)
	traverse(root.Right, result)
	*result = append(*result, root.Val)
}
