package _9_combination_sum

// NeetCode の解法見て書いたが...。むずいなこれは

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	var dfs func(idx int, cur []int, total int)
	dfs = func(idx int, cur []int, total int) {
		if total == target {
			res = append(res, copy(cur))
			return
		}
		if total > target || idx > len(candidates)-1 {
			return
		}
		// choose idx val
		cur = append(cur, candidates[idx])
		dfs(idx, cur, total+candidates[idx])

		// backtrack: 1個前に戻って次の数字を試す
		cur = cur[:len(cur)-1]
		dfs(idx+1, cur, total)
	}
	dfs(0, []int{}, 0)
	return res
}

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	path := []int{}

	var dfs func(int, int)
	dfs = func(i int, t int) {
		for i < len(candidates) {
			val := candidates[i]
			path = append(path, val)
			if t == val {
				// found
				res = append(res, copy(path))
			}
			if t > val {
				dfs(i, t-val)
			}
			// backtracking
			path = path[:len(path)-1]
			i++
		}
	}
	dfs(0, target)

	return res
}

func copy(orig []int) []int {
	cp := make([]int, len(orig))
	for i, v := range orig {
		cp[i] = v
	}
	return cp
}
