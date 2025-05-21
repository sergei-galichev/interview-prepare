package main

import (
	"math"
)

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	merged := mergeTwoLists(list1, list2)

	printList(merged)

}

func printList(list *ListNode) {
	for ; list != nil; list = list.Next {
		print(list.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
*
  - Definition for singly-linked list.
  - type ListNode struct {
  - Val int
  - Next *ListNode
  - }

time: O(n), n is node count
mem: O(1)
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyNode := &ListNode{0, nil}

	stab := dummyNode

	for list1 != nil || list2 != nil {
		if getNodeValue(list1) < getNodeValue(list2) {
			dummyNode.Next = list1
			list1 = list1.Next
		} else {
			dummyNode.Next = list2
			list2 = list2.Next
		}

		dummyNode = dummyNode.Next
	}

	return stab.Next
}

func getNodeValue(node *ListNode) float64 {
	if node == nil {
		return math.Inf(1)
	}

	return float64(node.Val)
}
