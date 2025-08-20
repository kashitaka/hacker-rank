package _15_kth_largest_element_in_an_array

import "container/heap"

type MaxHeap []int

func (m MaxHeap) Len() int           { return len(m) }
func (m MaxHeap) Less(i, j int) bool { return m[i] > m[j] }
func (m MaxHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &MaxHeap{}
	*h = append(*h, nums...)
	heap.Init(h)
	res := 0
	for i := 0; i < k; i++ {
		res = heap.Pop(h).(int)
	}
	return res
}
