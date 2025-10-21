package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))

}

/*
time: O(n * number of bits in the largest number), n is integer number
mem: O(1)
*/

func countBitsNonOptimized(n int) []int {
	result := make([]int, n+1)

	index := 0

	for i := 0; i <= n; i++ {
		binaryFormat := strconv.FormatInt(int64(i), 2)

		count := 0

		for _, v := range binaryFormat {
			bit, _ := strconv.Atoi(string(v))

			count += bit
		}

		result[index] += count

		index++
	}

	return result
}

func countBits(n int) []int {
	result := make([]int, n+1)

	for i := 0; i <= n; i++ {
		count := 0

		num := i

		for num > 0 {

			count += num & 1

			num >>= 1
		}

		result[i] = count
	}

	return result
}
