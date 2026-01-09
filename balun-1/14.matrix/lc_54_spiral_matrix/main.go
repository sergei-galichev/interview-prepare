package main

import "fmt"

func main() {
	m1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	msp := spiralOrder(m1)

	fmt.Println(msp)
}

func spiralOrder(matrix [][]int) []int {
	rowSize := len(matrix)
	colSize := len(matrix[0])

	top := 0
	bottom := rowSize - 1
	left := 0
	right := colSize - 1

	result := make([]int, 0, rowSize*colSize)

	for top <= bottom && left <= right {
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}

		top++

		if top > bottom {
			return result
		}

		for row := top; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}

		right--

		if left > right {
			return result
		}

		for col := right; col >= left; col-- {
			result = append(result, matrix[bottom][col])
		}

		bottom--

		for row := bottom; row >= top; row-- {
			result = append(result, matrix[row][left])
		}

		left++
	}

	return result
}
