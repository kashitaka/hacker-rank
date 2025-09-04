package _29_longest_increasing_path_in_a_matrix

func longestIncreasingPath(matrix [][]int) int {
	ROW, COL := len(matrix), len(matrix[0])
	dp := make([][]int, ROW)
	for i, _ := range dp {
		dp[i] = make([]int, COL)
	}

	var dfs func(r, c, prev int) int
	dfs = func(r, c, prev int) int {
		if r < 0 || c < 0 || r >= ROW || c >= COL ||
			matrix[r][c] <= prev {
			return 0
		}
		if dp[r][c] != 0 {
			return dp[r][c]
		}
		res := 0
		res = max(res, dfs(r-1, c, matrix[r][c]))
		res = max(res, dfs(r+1, c, matrix[r][c]))
		res = max(res, dfs(r, c-1, matrix[r][c]))
		res = max(res, dfs(r, c+1, matrix[r][c]))
		dp[r][c] = res + 1
		return res + 1
	}
	res := 0
	for r, row := range matrix {
		for c, _ := range row {
			res = max(res, dfs(r, c, -1))
		}
	}
	return res
}
