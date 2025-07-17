package _17_pacific_atlantic_water_flow

// Neet Code の解説動画見て、とりあえず自分で書いてみる。
// これは難しいねーーー

func pacificAtlantic(heights [][]int) [][]int {
	ROWS, COLS := len(heights), len(heights[0])
	pac, atl := make(map[[2]int]bool), make(map[[2]int]bool)

	var dfs func(row int, col int, visited map[[2]int]bool, prevHeight int)
	dfs = func(row int, col int, visited map[[2]int]bool, prevHeight int) {
		if row < 0 || col < 0 || row == ROWS || col == COLS ||
			visited[[2]int{row, col}] || // already visited
			heights[row][col] < prevHeight { // can't go lower place
			return
		}
		// this grid is reachable from the ocean
		visited[[2]int{row, col}] = true
		// go next
		dfs(row-1, col, visited, heights[row][col])
		dfs(row+1, col, visited, heights[row][col])
		dfs(row, col-1, visited, heights[row][col])
		dfs(row, col+1, visited, heights[row][col])
	}

	for c := 0; c < COLS; c++ {
		dfs(0, c, pac, heights[0][c])           // the most top row -> to the pac
		dfs(ROWS-1, c, atl, heights[ROWS-1][c]) // the most bottom row -> to the atl
	}

	for r := 0; r < ROWS; r++ {
		dfs(r, 0, pac, heights[r][0])           // the most left col -> to the pac
		dfs(r, COLS-1, atl, heights[r][COLS-1]) // the most bottom row -> to the atl
	}

	res := [][]int{}
	for row, v := range heights {
		for col, _ := range v {
			if pac[[2]int{row, col}] && atl[[2]int{row, col}] {
				res = append(res, []int{row, col})
			}
		}
	}
	return res
}
