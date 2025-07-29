package _6_merge_intervals

import "sort"

func merge(intervals [][]int) [][]int {
	// O(n log n)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	prev := intervals[0]

	// O(n)
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] >= cur[0] {
			// prev is overlapping. so marge
			prev[0] = min(prev[0], cur[0])
			prev[1] = max(prev[1], cur[1])
			continue
		}
		res = append(res, prev)
		prev = cur
	}
	res = append(res, prev)
	return res
}
