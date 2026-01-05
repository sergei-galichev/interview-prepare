[867. Transpose Matrix](https://leetcode.com/problems/transpose-matrix/)

```go
package main

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
```

***Оценка по времени:*** `O(n x m)`

*Объяснения:* фактически нужно перебрать все элементы матрицы размерностью `n x m` 

***Оценка по памяти:*** `O(1)`

*Объяснения:* если не учитывать результирующую матрицу

**Описание решения**

транспонирование матрицы - строки становятся столбцами, а столбцы - строками.

фиксируем количество строк и столбцов. создаем результирующую матрицу. идем по всем столбцам матрицы в цикле и создаем
для каждого столбца слайс размером на количество строк. идем в цикле по строкам и записываем в элемент
`result[столбец][строка]` результирующей матрицы значение `matrix[строка][столбец]` из исходной матрицы
