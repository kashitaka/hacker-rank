package _10_course_schedule_2

func findOrder(numCourses int, prerequisites [][]int) []int {
	prereq := make(map[int][]int)
	for _, v := range prerequisites {
		prereq[v[0]] = append(prereq[v[0]], v[1])
	}
	res := []int{}
	visited, visiting := make(map[int]bool), make(map[int]bool)

	var dfs func(crs int) bool
	dfs = func(crs int) bool {
		if visited[crs] {
			// already finished dfs search! return TRUE!
			return true
		}
		if visiting[crs] {
			// cycle detected!! return FALSE
			return false
		}
		// first time to visit this crs so far
		// mark crs as visiting, so that possible to detect cycle later.
		visiting[crs] = true
		adjs := prereq[crs]
		for _, nextCrs := range adjs {
			if !dfs(nextCrs) {
				return false
			}
		}
		// this crs can be finished! mark it visited and add to output
		visited[crs] = true
		res = append(res, crs)
		// after finishing dfs, we mark this crs not visiting
		visiting[crs] = false
		return true
	}
	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return []int{}
		}
	}
	return res
}
