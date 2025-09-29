package _73_k_closest_points_to_origin

import "container/heap"

type maxHeap [][]int

func (m maxHeap) Len() int { return len(m) }
func (m maxHeap) Less(i, j int) bool {
	return m[i][0]*m[i][0]+m[i][1]*m[i][1] > m[j][0]*m[j][0]+m[j][1]*m[j][1]
}
func (m maxHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func (m *maxHeap) Push(x any) {
	*m = append(*m, x.([]int))
}
func (m *maxHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := &maxHeap{}
	*h = append(*h, points...)
	heap.Init(h)
	for h.Len() > k {
		heap.Pop(h)
	}
	return *h
}
