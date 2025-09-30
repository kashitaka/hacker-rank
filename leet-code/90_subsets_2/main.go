package _0_subsets_2

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	res := [][]int{}
	path := []int{}
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(nums) {
			cp := append([]int{}, path...)
			res = append(res, cp)
			return
		}
		path = append(path, nums[i])
		dfs(i + 1)
		// backtracking and skip duplicates
		path = path[:len(path)-1]
		for i+1 < len(nums) && nums[i+1] == nums[i] {
			i++
		}
		dfs(i + 1)
	}
	dfs(0)
	return res
}
