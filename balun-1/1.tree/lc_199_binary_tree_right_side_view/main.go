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
			Val: 3,
			//Right: &TreeNode{
			//	Val: 4,
			//},
		},
	}

	fmt.Println(rightSideView(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
time: O(n), n is node count
mem: O(n) n is node count
*/

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)

	levelOrder(root, 0, &result)

	return result
}

func levelOrder(root *TreeNode, level int, result *[]int) {
	if root == nil {
		return
	}

	if level == len(*result) {
		*result = append(*result, 0)
	}

	(*result)[level] = root.Val
	levelOrder(root.Left, level+1, result)
	levelOrder(root.Right, level+1, result)
}
