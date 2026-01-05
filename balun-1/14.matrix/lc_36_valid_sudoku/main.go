package main

func main() {}

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
