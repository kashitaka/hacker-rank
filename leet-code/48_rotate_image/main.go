package _8_rotate_image

func rotate(matrix [][]int) {
	l, r := 0, len(matrix)-1

	for l < r {
		for i := 0; i < r-l; i++ {
			top, bottom := l, r

			topLeft := matrix[top][l+i]
			matrix[top][l+i] = matrix[bottom-i][l]
			matrix[bottom-i][l] = matrix[bottom][r-i]
			matrix[bottom][r-i] = matrix[top+i][r]
			matrix[top+i][r] = topLeft
		}
		l++
		r--
	}
}

func rotateNoInPlace(matrix [][]int) {
	WIDTH := len(matrix)
	res := make([][]int, WIDTH)
	for i, _ := range res {
		res[i] = make([]int, WIDTH)
	}

	for r, row := range matrix {
		for c, _ := range row {
			res[c][WIDTH-1-r] = matrix[r][c]
		}
	}
	copy(matrix, res)
}
