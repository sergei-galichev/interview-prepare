package main

import "fmt"

func main() {
	m1 := []int{1, 2, 3, 0, 0, 0}
	m2 := []int{2, 5, 6}

	merge(m1, 3, m2, 3)

	fmt.Println(m1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	first := m - 1
	second := n - 1

	for idx := m + n - 1; idx >= 0; idx-- {
		if second < 0 {
			break
		}

		if first >= 0 && nums1[first] >= nums2[second] {
			nums1[idx] = nums1[first]

			first--
		} else {
			nums1[idx] = nums2[second]

			second--
		}
	}
}
