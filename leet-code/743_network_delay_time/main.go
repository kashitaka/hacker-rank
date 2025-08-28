package _43_network_delay_time

import "container/heap"

type edge struct {
	node int
	dist int
}

type minHeap []*edge

func (m minHeap) Len() int { return len(m) }
func (m minHeap) Less(i, j int) bool {
	return m[i].dist < m[j].dist
}
func (m minHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func (m *minHeap) Push(x any) {
	*m = append(*m, x.(*edge))
}

func (m *minHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func networkDelayTime(times [][]int, n int, k int) int {
	adj := make(map[int][]*edge)
	for _, v := range times {
		adj[v[0]] = append(adj[v[0]], &edge{
			node: v[1],
			dist: v[2],
		})
	}
	visit := make(map[int]bool)
	h := &minHeap{&edge{
		node: k, dist: 0,
	}}
	heap.Init(h)
	t := 0
	for h.Len() > 0 {
		item := heap.Pop(h).(*edge)
		if visit[item.node] {
			continue
		}
		visit[item.node] = true
		t = max(t, item.dist)

		for _, next := range adj[item.node] {
			if visit[next.node] {
				continue
			}
			heap.Push(h, &edge{
				node: next.node,
				dist: item.dist + next.dist,
			})
		}
	}

	if len(visit) != n {
		return -1
	}
	return t
}
