package _00_number_of_island

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	checked := make([][]bool, m)
	cnt := 0
	for i := 0; i < m; i++ {
		row := make([]bool, n)
		checked[i] = row
	}
	for i, row := range grid {
		for j, cell := range row {
			if checked[i][j] == true {
				continue
			}
			if cell == 49 && !checked[i][j] {
				cnt++
				dfs(grid, i, j, checked)
			}
			checked[i][j] = true
		}
	}
	return cnt
}

func dfs(grid [][]byte, i, j int, checked [][]bool)  {
	if checked[i][j] {
		return
	}
	checked[i][j] = true
	if grid[i][j] == 48 {
		return
	}
	if i != 0 {
		dfs(grid, i-1, j, checked)
	}
	if i != len(grid)-1 {
		dfs(grid, i+1, j, checked)
	}
	if j != 0 {
		dfs(grid, i, j-1, checked)
	}
	if j != len(grid[0])-1 {
		dfs(grid, i, j+1, checked)
	}
}