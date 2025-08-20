package _8_subsets

func subsets(nums []int) [][]int {
	res := [][]int{[]int{}}
	path := []int{}
	var dfs func(int)
	dfs = func(i int) {
		for j := i; j < len(nums); j++ {
			path = append(path, nums[j])
			cp := make([]int, len(path))
			copy(cp, path)
			res = append(res, cp)
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}
