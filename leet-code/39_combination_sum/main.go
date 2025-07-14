package _9_combination_sum

// NeetCode の解法見て書いたが...。むずいなこれは

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	var dfs func(idx int, cur []int, total int)
	dfs = func(idx int, cur []int, total int) {
		if total == target {
			curCopy := make([]int, len(cur))
			copy(curCopy, cur)
			res = append(res, curCopy)
			return
		}
		if total > target || idx > len(candidates)-1 {
			return
		}
		// choose idx val
		cur = append(cur, candidates[idx])
		dfs(idx, cur, total+candidates[idx])

		// no choose idx val
		cur = cur[:len(cur)-1]
		dfs(idx+1, cur, total)
	}
	dfs(0, []int{}, 0)
	return res
}
