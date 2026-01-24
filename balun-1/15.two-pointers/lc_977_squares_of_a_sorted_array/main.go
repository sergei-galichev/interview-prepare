package main

import (
	"fmt"
)

func main() {
	nums := []int{-4, -1, 0, 3, 10}

	result := sortedSquares(nums)

	fmt.Println(result)
}

func sortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1

	result := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		lhs := abs(nums[left])  // math.Abs
		rhs := abs(nums[right]) // math.Abs

		if lhs >= rhs {
			result[i] = lhs * lhs
			left++
		} else {
			result[i] = rhs * rhs
			right--
		}
	}

	return result
}

func abs(number int) int {
	if number < 0 {
		return -number
	} else {
		return number
	}
}

//func sortedSquares(nums []int) []int {
//	var result []int
//
//	for i := 0; i < len(nums); i++ {
//		result = append(result, nums[i]*nums[i])
//	}
//
//	slices.Sort(result)
//
//	return result
//}
