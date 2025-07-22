package _23_number_of_connected_components

// Time: O(N)
// Space: O(N)
func countComponents(n int, edges [][]int) int {
	nodeMap := make(map[int][]int, n)
	// initialize node map
	// Time: O(N)
	// Space: O(N)
	for i := 0; i < n; i++ {
		nodeMap[i] = []int{}
	}

	for _, edge := range edges {
		n1, n2 := edge[0], edge[1]
		nodeMap[n1] = append(nodeMap[n1], n2)
		nodeMap[n2] = append(nodeMap[n2], n1)
	}
	visited := make(map[int]bool)
	res := 0

	var dfs func(int)
	dfs = func(node int) {
		if visited[node] {
			return
		}
		visited[node] = true
		// visit neighboring nodes
		for _, v := range nodeMap[node] {
			dfs(v)
		}
	}

	// Visit all nodes
	// Time: O(N)
	// Space: O(N) for recursive call
	for nodeVal, _ := range nodeMap {
		if !visited[nodeVal] {
			res++
			dfs(nodeVal)
		}
	}
	return res
}
