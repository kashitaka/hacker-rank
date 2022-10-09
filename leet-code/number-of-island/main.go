package main

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	seen := make([][]bool, m)
	for i := 0; i < m; i++ {
		row := make([]bool, n)
		for j := 0; j < n; j++ {
			row[j] = false
		}
		seen[i] = row
	}
	res := 0
	// 探索
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !seen[i][j] && grid[i][j] == 49 {
				res += 1
				// i,jから探索開始
				dfs(grid, seen, i, j, m, n)
			} else {
				seen[i][j] = true
			}
		}
	}
	return res
}

func dfs(grid [][]byte, seen [][]bool, i, j, m, n int) {
	if seen[i][j] {
		return
	}
	seen[i][j] = true
	if grid[i][j] == 48 {
		return
	}
	// 到達可能な部分を探索
	// 上のマス
	if i != 0 {
		dfs(grid, seen, i-1, j, m, n)
	}
	// 左のマス
	if j != 0 {
		dfs(grid, seen, i, j-1, m, n)
	}
	// 右のマス
	if j != n-1 {
		dfs(grid, seen, i, j+1, m, n)
	}
	// 下のマス
	if i != m-1 {
		dfs(grid, seen, i+1, j, m, n)
	}
}
