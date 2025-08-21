package _0_combination_sum_2

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	res := [][]int{}
	var dfs func(idx int, cur []int, t int)
	dfs = func(idx int, cur []int, t int) {
		if t == 0 {
			cp := make([]int, len(cur))
			copy(cp, cur)
			res = append(res, cp)
			return
		}
		if t < 0 || idx >= len(candidates) {
			return
		}
		cur = append(cur, candidates[idx])
		dfs(idx+1, cur, t-candidates[idx])
		// backtrack
		cur = cur[:len(cur)-1]

		// skip pattern
		for idx+1 < len(candidates) && candidates[idx] == candidates[idx+1] {
			idx++
		}
		dfs(idx+1, cur, t)
	}
	dfs(0, []int{}, target)
	return res
}
