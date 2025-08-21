package _9_word_search

func exist(board [][]byte, word string) bool {
	ROW, COL := len(board), len(board[0])
	visit := make([][]bool, ROW)
	for i, _ := range visit {
		visit[i] = make([]bool, COL)
	}

	var dfs func(row, col, idx int) bool
	dfs = func(row, col, idx int) bool {
		if row < 0 || col < 0 || row == ROW || col == COL ||
			visit[row][col] || word[idx] != board[row][col] {
			// this is wrong path
			return false
		}
		if idx == len(word)-1 {
			// found!!
			return true
		}
		// this maybe right path. mark this cell visited and keep searching adjuscent cells
		visit[row][col] = true
		res := dfs(row-1, col, idx+1) || dfs(row+1, col, idx+1) ||
			dfs(row, col-1, idx+1) || dfs(row, col+1, idx+1)
		// backtrack
		visit[row][col] = false
		return res
	}

	for r, row := range board {
		for c, _ := range row {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}
