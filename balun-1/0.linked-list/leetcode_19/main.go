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
memory: O(1)
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{0, head}

	length := 0

	curr := dummyNode

	for curr != nil {
		curr = curr.Next

		length++
	}

	curr = dummyNode

	for i := 0; i < length-n-1; i++ {
		curr = curr.Next
	}

	curr.Next = curr.Next.Next

	return dummyNode.Next
}
