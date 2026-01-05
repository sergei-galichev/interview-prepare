package main

import "fmt"

func main() {

	m1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	mtr := transpose(m1)

	fmt.Println(mtr)
}

func transpose(matrix [][]int) [][]int {
	rowSize := len(matrix)
	colSize := len(matrix[0])

	result := make([][]int, colSize)

	for column := 0; column < colSize; column++ {
		result[column] = make([]int, rowSize)

		for row := 0; row < rowSize; row++ {
			result[column][row] = matrix[row][column]
		}
	}

	return result
}
