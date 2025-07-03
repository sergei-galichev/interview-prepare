package main

import (
	"fmt"
	"slices"
)

func main() {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(levelOrderBottom(tree))
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

func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)

	postorder(root, 0, &result)

	return result
}

func postorder(root *TreeNode, level int, result *[][]int) {
	if root == nil {
		return
	}

	if level == len(*result) {
		*result = slices.Insert(*result, 0, make([]int, 0))
	}

	postorder(root.Left, level+1, result)
	postorder(root.Right, level+1, result)

	(*result)[len(*result)-level-1] = append((*result)[len(*result)-level-1], root.Val)
}
