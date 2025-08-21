package _31_palindrome_partitioning

func partition(s string) [][]string {
	res := [][]string{}
	path := []string{}

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(s) {
			cp := append([]string{}, path...)
			res = append(res, cp)
			return
		}
		for j := idx + 1; j <= len(s); j++ {
			sub := s[idx:j]
			if !isPalindrome(sub) {
				continue
			}
			path = append(path, sub)
			dfs(j)
			// backtrack
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
