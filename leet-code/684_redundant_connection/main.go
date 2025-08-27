package _84_redundant_connection

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	rank := make([]int, n+1)

	for i, _ := range parent {
		parent[i] = i
		rank[i] = 1
	}

	// find returns parent of given node
	// this updates parent slice when it traverses node's parents
	var find func(node int) int
	find = func(node int) int {
		if parent[node] != node {
			// may need to update.
			// path compression
			// ここが一番イメージ湧きにくい！
			parent[node] = find(parent[node])
		}
		return parent[node]
	}

	// union return true if node X and Y are already in the same root group by joining the 2 nodes to 1 group.
	var union func(x, y int) bool
	union = func(x, y int) bool {
		parentX, parentY := find(x), find(y)
		if parentX == parentY {
			// cycle detected
			return true
		}
		// otherwise, it needs to join 2 group into a single group
		// note that x and y may not be directly modified.
		// this modifies their parents' parent
		if rank[parentX] > rank[parentY] {
			// parent Y should belong to parent X
			parent[parentY] = parentX
			// and parentX's rank gets stronger
			rank[parentX] += rank[parentY]
		} else {
			// vice versa
			parent[parentX] = parentY
			rank[parentY] += rank[parentX]
		}
		return false
	}

	for _, edge := range edges {
		if union(edge[0], edge[1]) {
			return edge
		}
	}
	return []int{}
}
