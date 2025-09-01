package _87_cheapest_flights_within_k_stops

import (
	"container/heap"
	"math"
)

type node struct {
	n    int
	cost int
	stop int
}

type minHeap []*node

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i].cost < m[j].cost }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *minHeap) Push(x any) {
	*m = append(*m, x.(*node))
}
func (m *minHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

type flight struct {
	n    int
	cost int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	adj := make(map[int][]*flight)
	for i := 0; i < n; i++ {
		adj[i] = []*flight{}
	}
	for _, v := range flights {
		adj[v[0]] = append(adj[v[0]], &flight{
			n:    v[1],
			cost: v[2],
		})
	}

	h := &minHeap{&node{n: src, cost: 0, stop: 0}}
	heap.Init(h)
	res := math.MaxInt32
	visit := make(map[int]int) // node vs cost
	for h.Len() > 0 {
		pop := heap.Pop(h).(*node)
		if pop.n == dst {
			res = min(res, pop.cost)
			continue
		}
		if pop.stop > k {
			continue
		}
		if prevStop, ok := visit[pop.n]; ok && prevStop <= pop.stop {
			// このノードに来たのが2回目以降かつ、今回のstopの方が多い。
			// (costは最小順に到達することがminHeapで保証されてるので管理不要)
			continue
		}
		visit[pop.n] = pop.stop
		for _, next := range adj[pop.n] {
			if prevStop, ok := visit[next.n]; ok && prevStop <= pop.stop+1 {
				continue
			}
			heap.Push(h, &node{
				n:    next.n,
				cost: pop.cost + next.cost,
				stop: pop.stop + 1,
			})
		}
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}
