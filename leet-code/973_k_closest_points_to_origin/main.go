package _73_k_closest_points_to_origin

import "container/heap"

type MaxHeap [][]int

func (m MaxHeap) Len() int { return len(m) }
func (m MaxHeap) Less(i, j int) bool {
	distI := m[i][0]*m[i][0] + m[i][1]*m[i][1]
	distJ := m[j][0]*m[j][0] + m[j][1]*m[j][1]
	return distI < distJ
}
func (m MaxHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.([]int))
}
func (m *MaxHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := &MaxHeap{}
	for _, v := range points {
		*h = append(*h, v)
	}
	heap.Init(h)

	res := [][]int{}
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).([]int))
	}
	return res
}
