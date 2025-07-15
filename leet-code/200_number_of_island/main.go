package _00_number_of_island

func numIslands(grid [][]byte) int {
	rMax := len(grid) - 1
	cMax := len(grid[0]) - 1
	res := 0
	seen := make([][]bool, len(grid))
	for i, _ := range grid {
		seen[i] = make([]bool, len(grid[0]))
	}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r > rMax || c > cMax || seen[r][c] {
			return
		}
		seen[r][c] = true
		if grid[r][c] == '0' {
			return
		}
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for r, row := range grid {
		for c, _ := range row {
			if grid[r][c] == '1' && !seen[r][c] {
				res++
			}
			dfs(r, c)
		}
	}
	return res
}
