package course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
	}
	for i, _ := range graph {
		// iについての探索
		seen := make([]bool, len(graph))
		res := dfs(i, graph, seen, i)
		if !res {
			return false
		}
	}
	return true
}

func dfs(start int, graph [][]int, seen []bool, now int) bool {
	if seen[now] && now == start {
		return false
	}
	if seen[now] {
		return true
	}
	seen[now] = true
	nextVs := graph[now]
	for _, v := range nextVs {
		res := dfs(start, graph, seen, v)
		if !res {
			return false
		}
	}
	return true
}
