package _17_pacific_atlantic_water_flow

// Neet Code の解説動画見て、とりあえず自分で書いてみる。
// これは難しいねーーー
// 追記: pacとatl のキーを []int にせず [][] にすると高速になる。[]int だとpointerの検索になるから？

func pacificAtlantic(heights [][]int) [][]int {
	ROW := len(heights)
	COL := len(heights[0])

	pac := make([][]bool, ROW)
	atl := make([][]bool, ROW)
	for i, _ := range pac {
		pac[i] = make([]bool, COL)
		atl[i] = make([]bool, COL)
	}

	var dfs func(row int, col int, prevLevel int, visited [][]bool)
	dfs = func(row int, col int, prevLevel int, visited [][]bool) {
		if row < 0 || row >= ROW || col < 0 || col >= COL ||
			heights[row][col] < prevLevel || visited[row][col] {
			return
		}
		// this sell can be reached from the ocean
		visited[row][col] = true
		curLevel := heights[row][col]
		// keep scaning for adjuscent cells
		dfs(row-1, col, curLevel, visited)
		dfs(row+1, col, curLevel, visited)
		dfs(row, col-1, curLevel, visited)
		dfs(row, col+1, curLevel, visited)
	}

	// O(2 * n * m )
	for i := 0; i < COL; i++ {
		dfs(0, i, 0, pac)
		dfs(ROW-1, i, 0, atl)
	}
	// O(2 * n * m )
	for i := 0; i < ROW; i++ {
		dfs(i, 0, 0, pac)
		dfs(i, COL-1, 0, atl)
	}
	res := [][]int{}
	// O(n * m )
	for r, row := range pac {
		for c, cell := range row {
			if cell && atl[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}
	// Time: O( n * m )
	// Space: O(n * m )
	return res
}
