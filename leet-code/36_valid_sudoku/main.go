package _6_valid_sudoku

import "strconv"

func isValidSudoku(board [][]byte) bool {
	rowHash := make([][]bool, 9)
	colHash := make([][]bool, 9)
	chunkHash := make([][]bool, 9)

	for i, _ := range rowHash {
		rowHash[i] = make([]bool, 9)
		colHash[i] = make([]bool, 9)
		chunkHash[i] = make([]bool, 9)
	}

	for r, row := range board {
		for c, _ := range row {
			cell := board[r][c]
			if cell == '.' {
				continue
			}
			num, _ := strconv.Atoi(string(rune(cell)))
			num--
			chunckIdx := (r/3)*3 + c/3
			if rowHash[r][num] || colHash[c][num] || chunkHash[chunckIdx][num] {
				return false
			}
			rowHash[r][num] = true
			colHash[c][num] = true
			chunkHash[chunckIdx][num] = true
		}
	}
	return true
}
