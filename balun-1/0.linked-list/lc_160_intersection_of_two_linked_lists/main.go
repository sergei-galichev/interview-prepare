package main

func main() {
	common := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}

	list2 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  1,
			Next: common,
		},
	}

	list1 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  1,
				Next: common,
			},
		},
	}

	intersect := getIntersectionNode(list1, list2)

	printList(intersect)

}

func printList(list *ListNode) {
	for ; list != nil; list = list.Next {
		print(list.Val)
	}

	println()
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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	lengthA := listLength(headA)
	lengthB := listLength(headB)

	ptrA, ptrB := headA, headB

	if lengthA > lengthB {
		for i := 0; i < lengthA-lengthB; i++ {
			ptrA = ptrA.Next
		}
	} else if lengthA < lengthB {
		for i := 0; i < lengthB-lengthA; i++ {
			ptrB = ptrB.Next
		}
	}

	for ptrA != nil && ptrB != nil {
		if ptrA == ptrB && ptrA.Val == ptrB.Val {
			return ptrA
		}

		ptrA = ptrA.Next
		ptrB = ptrB.Next
	}

	return nil
}

func listLength(head *ListNode) int {
	length := 0

	curr := head

	for curr != nil {
		curr = curr.Next

		length++
	}

	return length
}
