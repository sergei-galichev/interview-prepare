package main

import "fmt"

func main() {

	m1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	sum1 := diagonalSum(m1)

	fmt.Printf("sum 1: %d\n", sum1)

	m2 := [][]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}

	sum2 := diagonalSum(m2)

	fmt.Printf("sum 2: %d\n", sum2)
}

func diagonalSum(mat [][]int) int {
	sum := 0

	length := len(mat)

	for i := 0; i < length; i++ {
		// add main diagonal element
		sum += mat[i][i]

		// add sub diagonal element
		sum += mat[length-1-i][i]
	}

	if length%2 != 0 {
		halfIdx := length / 2

		sum -= mat[halfIdx][halfIdx]
	}

	return sum
}
