package main

import "fmt"

func main() {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			//Left: &TreeNode{
			//	Val: 8,
			//},
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

	fmt.Println(zigzagLevelOrder(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
time: O(n), n is node count
mem: O(n) n is node count (O(2n))
*/

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)

	preorder(root, 0, &result)

	for i := 0; i < len(result); i++ {
		if i%2 == 1 {
			for j := 0; j < len(result[i])/2; j++ {
				result[i][j], result[i][len(result[i])-j-1] = result[i][len(result[i])-j-1], result[i][j]
			}
		}
	}

	return result
}

func preorder(root *TreeNode, level int, result *[][]int) {
	if root == nil {
		return
	}

	if level == len(*result) {
		*result = append(*result, make([]int, 0))
	}

	(*result)[level] = append((*result)[level], root.Val)

	preorder(root.Left, level+1, result)
	preorder(root.Right, level+1, result)
}
