package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15}
	//numbers := []int{2, 3, 4}

	result := twoSum(numbers, 9)
	//result := twoSum(numbers, 6)

	fmt.Println(result)
}

func twoSum(numbers []int, target int) []int {
	var result []int

	left := 0
	right := len(numbers) - 1

	for left != right {
		if numbers[left]+numbers[right] == target {
			result = append(result, left+1, right+1)
			break
		}

		if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		}
	}

	return result
}

//func twoSum(numbers []int, target int) []int {
//	var result []int
//
//	for slow := 0; slow < len(numbers); slow++ {
//		for fast := slow + 1; fast < len(numbers); fast++ {
//			if numbers[slow]+numbers[fast] == target {
//				result = append(result, slow+1, fast+1)
//				break
//			}
//		}
//	}
//
//	return result
//}
