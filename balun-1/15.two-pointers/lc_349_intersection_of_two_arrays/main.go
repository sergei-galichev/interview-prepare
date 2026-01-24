package main

import (
	"fmt"
	"slices"
)

func main() {
	//nums1 := []int{1, 2, 2, 1}
	nums1 := []int{4, 9, 5}
	//nums2 := []int{2, 2}
	nums2 := []int{9, 4, 9, 8, 4}

	result := intersection(nums1, nums2)

	fmt.Println(result)
}

func intersection(nums1 []int, nums2 []int) []int {
	slices.Sort(nums1)
	slices.Sort(nums2)

	idx1 := 0
	idx2 := 0

	var result []int

	for idx1 < len(nums1) && idx2 < len(nums2) {
		if nums1[idx1] < nums2[idx2] {
			idx1++
		} else if nums1[idx1] > nums2[idx2] {
			idx2++
		} else {
			result = append(result, nums1[idx1])

			for idx1 < len(nums1) && nums1[idx1] == result[len(result)-1] {
				idx1++
			}

			for idx2 < len(nums2) && nums2[idx2] == result[len(result)-1] {
				idx2++
			}
		}
	}

	return result
}

//func intersection(nums1 []int, nums2 []int) []int {
//	result := make([]int, 0)
//
//	for i := 0; i < len(nums1); i++ {
//		for j := 0; j < len(nums2); j++ {
//			if nums1[i] == nums2[j] {
//				result = append(result, nums1[i])
//				break
//			}
//		}
//	}
//
//	return result
//}
