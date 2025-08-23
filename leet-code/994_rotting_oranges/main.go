package _94_rotting_oranges

func orangesRotting(grid [][]int) int {
	ROW, COL := len(grid), len(grid[0])
	queue := [][]int{}

	// add queue
	for r, row := range grid {
		for c, _ := range row {
			if grid[r][c] == 2 {
				queue = append(queue, []int{r, c})
			}
		}
	}

	// loop until queue becomes zero
	i := -1
	for len(queue) > 0 {
		for _, item := range queue {
			queue = queue[1:]
			r, c := item[0], item[1]
			dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
			for _, dir := range dirs {
				newR, newC := r+dir[0], c+dir[1]
				if newR < 0 || newR == ROW || newC < 0 || newC == COL ||
					grid[newR][newC] != 1 {
					continue
				}
				grid[newR][newC] = 2
				queue = append(queue, []int{newR, newC})
			}
		}
		i++
	}

	// check wheter threre is fresh orange. if so, return -1
	for r, row := range grid {
		for c, _ := range row {
			if grid[r][c] == 1 {
				return -1
			}
		}
	}
	if i == -1 {
		return 0
	}
	return i
}
