package main

const (
	EMPTY int = 0
)

type Direction int

const (
	UP        Direction = 0
	DOWN      Direction = 1
	LEFT      Direction = 2
	RIGHT     Direction = 3
	UPLEFT    Direction = 4
	DOWNLEFT  Direction = 5
	UPRIGHT   Direction = 6
	DOWNRIGHT Direction = 7
)

type TicTacToe struct {
	grid [][]int
}

func Constructor(n int) TicTacToe {
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		row := make([]int, n)
		for j := 0; j < n; j++ {
			row[j] = EMPTY
		}
		grid[i] = row
	}
	return TicTacToe{grid: grid}
}

func (this *TicTacToe) Move(row int, col int, player int) int {
	this.grid[row][col] = player
	count := make([]int, 8)
	for i := 0; i <= 7; i++ {
		count[i] = this.countRecursively(row, col, player, Direction(i), 0)
	}
	if (count[UP]+count[DOWN]-1 == len(this.grid)) ||
		(count[LEFT]+count[RIGHT]-1 == len(this.grid)) ||
		(count[UPLEFT]+count[DOWNRIGHT]-1 == len(this.grid)) ||
		(count[UPRIGHT]+count[DOWNLEFT]-1 == len(this.grid)) {
		return player
	}
	return 0
}

func (this *TicTacToe) countRecursively(row int, col int, player int, direction Direction, res int) int {
	if this.grid[row][col] != player {
		return res
	}
	res++
	nextRow, nextCol := 0, 0
	switch direction {
	case DOWN:
		if row == len(this.grid)-1 {
			return res
		}
		nextRow = row + 1
		nextCol = col
	case UP:
		if row == 0 {
			return res
		}
		nextRow = row - 1
		nextCol = col
	case LEFT:
		if col == 0 {
			return res
		}
		nextRow = row
		nextCol = col - 1
	case RIGHT:
		if col == len(this.grid)-1 {
			return res
		}
		nextRow = row
		nextCol = col + 1
	case UPLEFT:
		if row == 0 || col == 0 {
			return res
		}
		nextRow = row - 1
		nextCol = col - 1
	case UPRIGHT:
		if row == 0 || col == len(this.grid)-1 {
			return res
		}
		nextRow = row - 1
		nextCol = col + 1
	case DOWNLEFT:
		if row == len(this.grid)-1 || col == 0 {
			return res
		}
		nextRow = row + 1
		nextCol = col - 1
	case DOWNRIGHT:
		if row == len(this.grid)-1 || col == len(this.grid)-1 {
			return res
		}
		nextRow = row + 1
		nextCol = col + 1
	}
	return this.countRecursively(nextRow, nextCol, player, direction, res)
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
