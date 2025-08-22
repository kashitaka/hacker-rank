package _95_max_area_of_island

func maxAreaOfIsland(grid [][]int) int {
	ROW, COL := len(grid), len(grid[0])
	visit := make([][]bool, ROW)
	for i, _ := range visit {
		visit[i] = make([]bool, COL)
	}
	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || c < 0 || r == ROW || c == COL ||
			visit[r][c] || grid[r][c] == 0 {
			return 0
		}
		visit[r][c] = true
		return 1 + dfs(r-1, c) + dfs(r+1, c) + dfs(r, c-1) + dfs(r, c+1)
	}

	res := 0
	for r, row := range grid {
		for c, _ := range row {
			res = max(res, dfs(r, c))
		}
	}
	return res
}
