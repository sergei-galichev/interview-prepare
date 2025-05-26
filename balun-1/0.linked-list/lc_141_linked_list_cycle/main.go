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

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

func hasCycleBad(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	visited := make(map[*ListNode]struct{})

	cur := head

	for cur != nil {
		if _, ok := visited[cur]; ok {
			return true
		}

		visited[cur] = struct{}{}

		cur = cur.Next
	}

	return false
}
