package _3_set_matrix_zeroes

// TIME: O(n * m)
// SPACE: O(n * m)
func setZeroes(matrix [][]int) {
	ROW := len(matrix)
	COL := len(matrix[0])
	changed := make([][]bool, ROW)
	for i, _ := range changed {
		changed[i] = make([]bool, COL)
	}

	var change func(row, col int)
	change = func(row, col int) {
		// right
		for i := col + 1; i < COL; i++ {
			if matrix[row][i] != 0 {
				matrix[row][i] = 0
				changed[row][i] = true
			}
		}

		// bottom
		for i := row + 1; i < ROW; i++ {
			if matrix[i][col] != 0 {
				matrix[i][col] = 0
				changed[i][col] = true
			}
		}

		// left
		for i := col - 1; i >= 0; i-- {
			if matrix[row][i] != 0 {
				matrix[row][i] = 0
				changed[row][i] = true
			}
		}

		// top
		for i := row - 1; i >= 0; i-- {
			if matrix[i][col] != 0 {
				matrix[i][col] = 0
				changed[i][col] = true
			}
		}
	}

	for r, row := range matrix {
		for c, _ := range row {
			if matrix[r][c] == 0 && !changed[r][c] {
				change(r, c)
			}
		}
	}
}
