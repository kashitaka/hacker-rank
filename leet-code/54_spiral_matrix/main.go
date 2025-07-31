package _4_spiral_matrix

func spiralOrderDFS(matrix [][]int) []int {
	ROW, COL := len(matrix), len(matrix[0])
	visit := make([][]bool, ROW)
	for i, _ := range visit {
		visit[i] = make([]bool, COL)
	}

	res := []int{}

	// direction
	// 0: go right / 1: go down / 2: go left / 3 go top
	var dfs func(r, c, direction int) bool
	dfs = func(r, c, direction int) bool {
		if r < 0 || r == ROW || c < 0 || c == COL ||
			visit[r][c] {
			return false
		}
		visit[r][c] = true
		res = append(res, matrix[r][c])

		// next move
		var goNext bool
		switch direction {
		case 0:
			goNext = dfs(r, c+1, 0)
		case 1:
			goNext = dfs(r+1, c, 1)
		case 2:
			goNext = dfs(r, c-1, 2)
		case 3:
			goNext = dfs(r-1, c, 3)
		}
		if goNext {
			return true
		}
		switch direction {
		case 0:
			goNext = dfs(r+1, c, 1)
		case 1:
			goNext = dfs(r, c-1, 2)
		case 2:
			goNext = dfs(r-1, c, 3)
		case 3:
			goNext = dfs(r, c+1, 0)
		}
		return true
	}

	dfs(0, 0, 0)
	return res
}
