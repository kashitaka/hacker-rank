package _30_surrounded_regions

func solve(board [][]byte) {
	ROW, COL := len(board), len(board[0])
	edge := make([][]bool, ROW) // manages whether each cell can go to edge
	for i, _ := range edge {
		edge[i] = make([]bool, COL)
	}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r == ROW || c == COL ||
			board[r][c] == 'X' || edge[r][c] {
			return
		}
		edge[r][c] = true
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	for r := 0; r < ROW; r++ {
		if board[r][0] == 'O' {
			dfs(r, 0)
		}
		if board[r][COL-1] == 'O' {
			dfs(r, COL-1)
		}
	}
	for c := 0; c < COL; c++ {
		if board[0][c] == 'O' {
			dfs(0, c)
		}
		if board[ROW-1][c] == 'O' {
			dfs(ROW-1, c)
		}
	}
	for r, row := range board {
		for c, _ := range row {
			if !edge[r][c] {
				board[r][c] = 'X'
			}
		}
	}
}
