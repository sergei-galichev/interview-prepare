package main

import (
	"math"
)

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
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

	list3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 6,
		},
	}

	lists := []*ListNode{list1, list2, list3}

	mergeList := mergeKLists(lists)

	printList(mergeList)
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

func mergeKListsBad(lists []*ListNode) *ListNode {
	minNodeIdx := minNodeIndex(lists)

	if minNodeIdx == nil {
		return nil
	}

	curr := lists[*minNodeIdx]
	head := curr

	lists[*minNodeIdx] = curr.Next

	for {
		minNodeIdx = minNodeIndex(lists)

		if minNodeIdx == nil {
			curr.Next = nil

			break
		}

		curr.Next = lists[*minNodeIdx]
		curr = curr.Next
		lists[*minNodeIdx] = curr.Next
	}

	return head
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummyNode := &ListNode{}

	curr := dummyNode

	for {
		minNodeIdx := minNodeIndex(lists)

		if minNodeIdx == nil {
			curr.Next = nil

			break
		}

		curr.Next = lists[*minNodeIdx]
		curr = curr.Next
		lists[*minNodeIdx] = lists[*minNodeIdx].Next
	}

	return dummyNode.Next
}

func minNodeIndex(lists []*ListNode) *int {
	minListNodeVal, minListNodeIndex := math.Inf(1), math.Inf(1)

	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}

		nodeValue := float64(lists[i].Val)

		if nodeValue < minListNodeVal {
			minListNodeVal = nodeValue
			minListNodeIndex = float64(i)
		}
	}

	if math.IsInf(minListNodeVal, 1) {
		return nil
	}

	index := int(minListNodeIndex)

	return &index
}
