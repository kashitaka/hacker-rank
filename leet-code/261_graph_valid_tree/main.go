package _61_graph_valid_tree

func validTree(n int, edges [][]int) bool {
	routeMap := make(map[int][]int, n)
	for i := 0; i < n; i++ {
		routeMap[i] = []int{}
	}

	visited := make(map[int]bool)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		routeMap[a] = append(routeMap[a], b)
		routeMap[b] = append(routeMap[b], a)
	}

	var dfs func(int, int) bool
	dfs = func(cur int, prev int) bool {
		if visited[cur] {
			return false
		}
		visited[cur] = true
		for _, v := range routeMap[cur] {
			if v != prev {
				if !dfs(v, cur) {
					return false
				}
			}
		}
		return true
	}

	return dfs(0, -1) && len(visited) == len(routeMap)
}
