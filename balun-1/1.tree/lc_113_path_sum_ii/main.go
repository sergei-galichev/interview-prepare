package main

import (
	"fmt"
)

func main() {
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

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := make([][]int, 0)
	prefixSlice := make([]int, 0)

	checkHasSum(root, 0, targetSum, &prefixSlice, &result)

	return result
}

func checkHasSum(root *TreeNode, prefixSum, targetSum int, prefixSlice *[]int, result *[][]int) {
	if root == nil {
		return
	}

	*prefixSlice = append(*prefixSlice, root.Val)
	defer func() {
		*prefixSlice = (*prefixSlice)[:len(*prefixSlice)-1]
	}()

	if isLeaf(root) && prefixSum+root.Val == targetSum {
		*result = append(*result, append([]int{}, *prefixSlice...))

		return
	}

	checkHasSum(root.Left, prefixSum+root.Val, targetSum, prefixSlice, result)
	checkHasSum(root.Right, prefixSum+root.Val, targetSum, prefixSlice, result)
}

func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}
