package _9_word_search

func exist(board [][]byte, word string) bool {
	ROW := len(board)
	COL := len(board[0])
	visit := make([][]bool, ROW)
	for i, _ := range visit {
		visit[i] = make([]bool, COL)
	}

	res := false
	var dfs func(int, int, int)
	dfs = func(row, col, idx int) {
		target := word[idx]
		if row < 0 || row == ROW || col < 0 || col == COL ||
			visit[row][col] || target != board[row][col] {
			return
		}
		if idx == len(word)-1 {
			// found!!!
			res = true
			return
		}
		visit[row][col] = true
		// find next
		dfs(row-1, col, idx+1)
		dfs(row+1, col, idx+1)
		dfs(row, col-1, idx+1)
		dfs(row, col+1, idx+1)
		// backtracking
		visit[row][col] = false
	}

	for r, rows := range board {
		for c, _ := range rows {
			dfs(r, c, 0)
			if res {
				return true
			}
		}
	}
	return false
}
