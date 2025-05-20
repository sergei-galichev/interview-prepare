package main

func main() {
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

func reorderList(head *ListNode) {
	tmp := preMiddle(head)

	ptr2 := reverseList(tmp.Next)

	tmp.Next = nil

	ptr1 := head

	for ptr2 != nil {
		ptr1Next := ptr1.Next
		ptr1.Next = ptr2
		ptr1 = ptr2
		ptr2 = ptr1Next
	}
}

func preMiddle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	curr := head

	for curr != nil {
		tmp := curr

		curr = curr.Next

		tmp.Next = prev

		prev = tmp
	}

	return prev
}
