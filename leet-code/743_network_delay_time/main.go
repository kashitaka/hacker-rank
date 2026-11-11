package _43_network_delay_time

import "container/heap"

type item struct {
	node int
	cost int
}

type minHeap []*item

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i].cost < m[j].cost }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *minHeap) Push(x any) {
	*m = append(*m, x.(*item))
}

func (m *minHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

// Prims algorithm: finds the minimum cost to span a given graph.
// Let V, E be the number of vertices and edges of the given graph
// The running time is O((V+E)logV), since for each node and edge,
// push and pop operations of a priority queue take O(log V).
func networkDelayTime(times [][]int, n int, k int) int {
	// O(E)
	graph := make(map[int][][2]int)
	for _, edge := range times {
		src, dest, weight := edge[0], edge[1], edge[2]
		graph[src] = append(graph[src], [2]int{dest, weight})
	}

	visit := make(map[int]bool)

	res := -1

	h := &minHeap{&item{k, 0}}
	heap.Init(h)

	// O(V+E)
	for h.Len() > 0 {
		pop := heap.Pop(h).(*item)
		src, time := pop.node, pop.cost
		if visit[pop.node] {
			continue
		}

		// mark as visit
		visit[src] = true
		res = time

		dests := graph[src]
		for _, dest := range dests {
			des, totalTime := dest[0], time+dest[1]
			if visit[des] {
				continue
			}

			// add queue item
			heap.Push(h, &item{des, totalTime})
		}
	}

	// O(1)
	if len(visit) != n {
		return -1
	}
	return res
}
