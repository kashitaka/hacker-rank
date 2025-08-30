package _78_swim_in_rising_water

import "container/heap"

type item struct {
	r     int
	c     int
	level int
}

type MinHeap []*item

func (m MinHeap) Len() int           { return len(m) }
func (m MinHeap) Less(i, j int) bool { return m[i].level < m[j].level }
func (m MinHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(*item))
}

func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func swimInWater(grid [][]int) int {
	ROW, COL := len(grid), len(grid[0])
	visit := make([][]bool, ROW)
	for i, _ := range visit {
		visit[i] = make([]bool, COL)
	}

	h := &MinHeap{&item{r: 0, c: 0, level: grid[0][0]}}
	heap.Init(h)
	res := 0
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for {
		pop := heap.Pop(h).(*item)
		if visit[pop.r][pop.c] {
			continue
		}
		visit[pop.r][pop.c] = true
		res = max(res, grid[pop.r][pop.c])
		if pop.r == ROW-1 && pop.c == COL-1 {
			break
		}
		for _, dir := range dirs {
			newR, newC := pop.r+dir[0], pop.c+dir[1]
			if newR < 0 || newC < 0 || newR == ROW || newC == COL ||
				visit[newR][newC] {
				continue
			}
			heap.Push(h, &item{r: newR, c: newC, level: grid[newR][newC]})
		}
	}

	return res
}
