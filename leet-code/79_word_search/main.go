package _9_word_search

func exist(board [][]byte, word string) bool {
	rowMax, colMax := len(board), len(board[0])
	path := make(map[[2]int]bool)

	var dfs func(r, c, i int) bool
	dfs = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}

		if r < 0 || c < 0 || r > rowMax-1 || c > colMax-1 ||
			board[r][c] != word[i] || // the cell's char is not the one seeking for
			path[[2]int{r, c}] { // the sell is already visited before
			return false
		}

		// luckily, the cell is valid
		// now we add this cell to the path and search neighboring cells
		path[[2]int{r, c}] = true
		res := dfs(r+1, c, i+1) ||
			dfs(r-1, c, i+1) ||
			dfs(r, c+1, i+1) ||
			dfs(r, c-1, i+1)
		delete(path, [2]int{r, c})
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
