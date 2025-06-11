package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}

	numArray := Constructor(nums)
	fmt.Println(numArray.prefixSum)
	fmt.Println(numArray.SumRange(0, 2))
	fmt.Println(numArray.SumRange(2, 5))
	fmt.Println(numArray.SumRange(0, 5))
}

type NumArray struct {
	prefixSum []int
}

/*
time: O(n), n is array length
mem: O(n) n is array length
*/

func Constructor(nums []int) NumArray {
	prefixSum := []int{0}

	for _, num := range nums {
		prefixSum = append(prefixSum, prefixSum[len(prefixSum)-1]+num)
	}

	return NumArray{
		prefixSum: prefixSum,
	}
}

/*
time: O(1), substracting two number
mem: O(1) substracting two number
*/

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSum[right+1] - this.prefixSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
