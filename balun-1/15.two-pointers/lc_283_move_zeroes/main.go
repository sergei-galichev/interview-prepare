package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 1, 0, 3, 12}

	moveZeroes(nums)

	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	slow := 0

	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
