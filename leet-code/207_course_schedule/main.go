package _07_course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	dep := make(map[int][]int)
	for _, v := range prerequisites {
		dep[v[0]] = append(dep[v[0]], v[1])
	}
	visit := make(map[int]bool)

	var dfs func(i int) bool
	dfs = func(i int) bool {
		if visit[i] {
			return false
		}
		next, ok := dep[i]
		if !ok || len(next) == 0 {
			return true
		}
		visit[i] = true
		for _, v := range next {
			if !dfs(v) {
				return false
			}
		}
		dep[i] = []int{}
		visit[i] = false
		return true
	}
	for k, _ := range dep {
		if !dfs(k) {
			return false
		}
	}
	return true
}
