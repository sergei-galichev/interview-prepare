package main

import (
	"log"
)

func main() {
	//list1 := &ListNode{
	//	Val: 2,
	//	Next: &ListNode{
	//		Val: 1,
	//		Next: &ListNode{
	//			Val: 5,
	//		},
	//	},
	//}

	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		},
	}

	result := nextLargerNodes(list1)

	log.Println(result)
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

func nextLargerNodesBad(head *ListNode) []int {
	ptr1 := head
	ptr2 := head

	result := make([]int, 0)

	for {
		if ptr1.Next == nil {
			result = append(result, 0)

			break
		}

		if ptr2 == nil {
			result = append(result, 0)

			tmp := ptr1.Next
			ptr1 = tmp
			ptr2 = tmp

			continue
		}

		if ptr2.Val > ptr1.Val {
			result = append(result, ptr2.Val)

			tmp := ptr1.Next
			ptr1 = tmp
			ptr2 = tmp

			continue
		}

		ptr2 = ptr2.Next
	}

	return result
}

func nextLargerNodes(head *ListNode) []int {
	nums := make([]int, 0)
	curr := head

	for curr != nil {
		nums = append(nums, curr.Val)
		curr = curr.Next
	}

	stack := make([]int, 0)
	result := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			result[i] = stack[len(stack)-1]
		}

		stack = append(stack, nums[i])
	}

	return result
}
