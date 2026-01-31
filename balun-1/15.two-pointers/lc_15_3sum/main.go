package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}

	result := threeSum(nums)

	fmt.Println(result)
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	var result [][]int

	for idx := 0; idx < len(nums)-2; idx++ {
		if idx > 0 && nums[idx] == nums[idx-1] {
			continue
		}

		left := idx + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[idx] + nums[left] + nums[right]

			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				result = append(result, []int{nums[idx], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return result
}
