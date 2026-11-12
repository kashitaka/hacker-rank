package _78_swim_in_rising_water

import "container/heap"

type MinHeap [][3]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][2] < h[j][2] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.([3]int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Let m x n be the size of the given grid.
// In the worst case, all cells may be visited, and each cell
// explores at most 4 directions, which costs O(4mn) = O(mn).
// Each heap push/pop operation takes O(log(mn)), and
// the total number of heap operations is O(mn).
// Therefore, the total running time is O(mn log(mn)).
func swimInWater(grid [][]int) int {
	ROW, COL := len(grid), len(grid[0])

	h := &MinHeap{{0, 0, grid[0][0]}}
	heap.Init(h)

	visit := make([][]bool, ROW)
	for i, _ := range visit {
		visit[i] = make([]bool, COL)
	}

	res := 0
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	for h.Len() > 0 {
		pop := heap.Pop(h).([3]int)
		r, c, level := pop[0], pop[1], pop[2]
		if visit[r][c] {
			continue
		}
		visit[r][c] = true
		res = max(res, level)

		if r == ROW-1 && c == COL-1 {
			break
		}

		for _, dir := range dirs {
			nextR, nextC := r+dir[0], c+dir[1]
			if nextR < 0 || nextC < 0 || nextR == ROW || nextC == COL ||
				visit[nextR][nextC] {
				continue
			}
			heap.Push(h, [3]int{nextR, nextC, grid[nextR][nextC]})
		}
	}

	return res
}
