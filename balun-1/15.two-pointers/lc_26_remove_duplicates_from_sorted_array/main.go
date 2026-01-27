package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}

	idx := removeDuplicates(nums)

	fmt.Println(idx)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0
	fast := 1

	for fast < len(nums) {
		if nums[slow] == nums[fast] {
			fast++
		} else if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
			fast++
		}
	}

	return slow + 1
}
