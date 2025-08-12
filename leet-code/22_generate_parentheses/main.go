package _2_generate_parentheses

// O(2*n)
func generateParenthesis(n int) []string {
	res := []string{}
	var dfs func(str string, openCnt int, closeCnt int)
	dfs = func(str string, openCnt int, closeCnt int) {
		if openCnt == n && closeCnt == n {
			res = append(res, str)
			return
		}
		if openCnt < n {
			dfs(str+"(", openCnt+1, closeCnt)
		}
		if closeCnt < openCnt {
			dfs(str+")", openCnt, closeCnt+1)
		}
	}
	dfs("", 0, 0)
	return res
}
