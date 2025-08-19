package _03_kth_largest_element_in_a_stream

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type KthLargest struct {
	minHeap *MinHeap
	Cap     int
}

func Constructor(k int, nums []int) KthLargest {
	h := &MinHeap{}
	heap.Init(h)
	res := KthLargest{
		minHeap: h,
		Cap:     k,
	}
	for _, v := range nums {
		res.Add(v)
	}
	return res
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.minHeap, val)
	if this.minHeap.Len() > this.Cap {
		heap.Pop(this.minHeap)
	}
	return (*this.minHeap)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
