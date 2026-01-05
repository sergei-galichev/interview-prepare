[36. Valid Sudoku](https://leetcode.com/problems/valid-sudoku/)

```go
package main

func isValidSudoku(board [][]byte) bool {
	var cols [9][9]bool
	var rows [9][9]bool
	var boxes [9][9]bool

	boxSize := 3
	boardSize := 9

	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			symbol := board[row][col]

			if symbol < '0' || symbol > '9' {
				continue
			}

			digit := symbol - '0' - 1

			if cols[col][digit] {
				return false
			}

			if rows[row][digit] {
				return false
			}

			boxRow := row / boxSize
			boxCol := col / boxSize

			box := boxRow*boxSize + boxCol

			if boxes[box][digit] {
				return false
			}

			cols[col][digit] = true
			rows[col][digit] = true
			boxes[col][digit] = true
		}

	}

	return true
}
```

***Оценка по времени:*** `O(1)`

*Объяснения:* размер входных данных фиксированный

***Оценка по памяти:*** `O(1)`

*Объяснения:* размер входных данных фиксированный

**Описание решения**

создаем массивы для хранения цифр в столбцах, строках и квадратах. добавляем переменные, содержащие размерность
квадратов и доски. идем в цикле по всем строкам. во внутреннем цикле идем по всем столбцам. получаем байтовое представ-
ление символа и проверяем, является ли он цифрой от 1 до 9. получаем цифру и проверяем встречается ли она в массивах
строк и столбцов. получаем промежуточные индексы, которые помогут найти индекс для массива квадратов. проверяем, что
цифра отсутствует в массиве. в массивы строк, столбцов и квадратов добавляем цифру 