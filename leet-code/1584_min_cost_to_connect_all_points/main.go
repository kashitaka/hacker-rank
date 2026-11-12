package _584_min_cost_to_connect_all_points

import "container/heap"

type item struct {
	node int
	cost int
}

type MinHeap []*item

func (h MinHeap) Len() int {return len(h)}
func (h MinHeap) Less(i, j int) bool {return h[i].cost < h[j].cost}
func (h MinHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(*item))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Minimum Spanning Tree with Prim's Algorithm
// For given graph G = (V, E), there are V^2 edges so |E| = |V|^2
// Need to check all nodes and edges so O(|V|+|E|) = O(|V|+|V|^2).
// In each iteration, push and pop operation needed for a priority queue.
// The maximum queue length might be number of edges. Thus, O(log |V|^2) =
// O(2 * log |V|) = O(log |V|) required.
// Hence, overall running time is O(|V|^2 * log(V))
func minCostConnectPoints(points [][]int) int {
	n := len(points)
	graph := make(map[int][2]int)
	for i, v := range points {
		graph[i] = [2]int{v[0],v[1]}
	}

	visit := make(map[int]bool)

	h := &MinHeap{&item{0, 0}}
	heap.Init(h)

	res := 0

	for len(visit) < n {
		pop := heap.Pop(h).(*item)
		node, cost := pop.node, pop.cost
		if visit[node] {
			continue
		}

		visit[node] = true
		res += cost

		srcX, srcY := graph[node][0], graph[node][1]
		for dest := 1; dest < n; dest++ {
			if visit[dest] || dest == node {
				continue
			}
			destX, destY := graph[dest][0], graph[dest][1]
			dist := abs(srcX-destX) + abs(srcY-destY)
			heap.Push(h, &item{dest, dist})
		}
	}

	return res
}