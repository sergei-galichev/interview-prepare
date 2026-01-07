package main

import "fmt"

func main() {
	m1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rotate(m1)

	fmt.Println(m1)
}

func rotate(matrix [][]int) {
	matrixSize := len(matrix)

	for row := 0; row < matrixSize; row++ {
		for col := row; col < matrixSize; col++ {
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}
	}

	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize/2; col++ {
			matrix[row][col], matrix[row][matrixSize-col-1] = matrix[row][matrixSize-col-1], matrix[row][col]
		}
	}
}
