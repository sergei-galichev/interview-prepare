[54. Spiral Matrix](https://leetcode.com/problems/spiral-matrix/)

```go
package main

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
```

***Оценка по времени:*** `O(n x m)`

*Объяснения:* проходим по всем элементам матрицы

***Оценка по памяти:*** `O(1)`

*Объяснения:* дополнительной памяти не выделяется

**Описание решения**

нужно завести 4 индекса: top, bottom, left, right. в цикле идем по элементам матрицы, начиная с top. постепенно двигаем
индексы по мере прохождения по элементам. сначала идем вправо от left и добавляем все элементы пока col <= right.
инкрементируем top и если он больше bottom, то сразу прерываем цикл. теперь идем вниз начиная с top и добавляем элементы
пока он меньше bottom. right уменьшаем на 1 и если left > right, то сразу прерываем цикл. теперь идем влево от right
пока col >= left. уменьшаем bottom на 1. и идем вверх от bottom пока row >= top. увеличиваем left на 1. цикл не преры-
вается пока top <= bottom и left <= right.