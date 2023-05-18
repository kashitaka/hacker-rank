package _046_last_stone_weight

import "container/heap"

func lastStoneWeight(stones []int) int {
	maxHeap := &MaxHeap{}
	for _, v := range stones {
		heap.Push(maxHeap, v)
	}

	for maxHeap.Len() > 1 {
		x := heap.Pop(maxHeap).(int)
		y := heap.Pop(maxHeap).(int)
		if x != y {
			heap.Push(maxHeap, x-y)
		}
	}

	if maxHeap.Len() == 0 {
		return 0
	}
	return (*maxHeap)[0]
}


type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Pop returns 最後を取り出して返す。元の配列は最後をカットしておく
func (h *MaxHeap) Pop() any {
	x := (*h)[len(*h)-1] // 最後の要素
	*h = (*h)[:len(*h)-1]
	return x
}
