[48. Rotate Image](https://leetcode.com/problems/rotate-image/)

```go
package main

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
```

***Оценка по времени:*** `O(n x n)`

*Объяснения:* обходим все элементы матрицы

***Оценка по памяти:*** `O(1)`

*Объяснения:* дополнительной памяти не выделяем

**Описание решения**

Во-первых, нужно транспонировать матрицу in-place. Во-вторых, поменять местами элементы матрицы в столбцах. Оба пункта
выполняем с помощью swap. В цикле обходим всю матрицу и свопаем элементы. На каждой новой строке текущий столбец устана-
вливается равным текущей строке. Так не свопаем элементы, в которых нет необходимости. Далеее, в цикле идем по матрице
и свопаем элементы в столбцах. Поэтому проходим только до среднего элемента