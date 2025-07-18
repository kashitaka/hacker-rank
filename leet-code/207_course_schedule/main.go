package _07_course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	preMap := make(map[int][]int)
	for i := 0; i < numCourses; i++ {
		preMap[i] = []int{}
	}

	for _, req := range prerequisites {
		preMap[req[0]] = append(preMap[req[0]], req[1])
	}

	visited := make(map[int]bool)

	var dfs func(crs int) bool
	dfs = func(crs int) bool {
		if visited[crs] {
			return false
		}
		if len(preMap[crs]) == 0 {
			return true
		}
		// ここの順番が大事。
		// len(preMap[crs]) == 0 って訪れとるやん
		// だが、preRequisitがないnodeは何度きてもいいので visitedフラグ立てない
		visited[crs] = true

		for _, nextCrs := range preMap[crs] {
			if !dfs(nextCrs) {
				return false
			}
		}
		// ここまできたってことはこのcrsは履修可能。
		visited[crs] = false  // ここでバックトラッキング的にフラグ寝かさないと再訪できない
		preMap[crs] = []int{} // 初期化
		return true
	}

	for k, _ := range preMap {
		if !dfs(k) {
			return false
		}
	}
	return true
}
