package _6_minimum_window_substring

func minWindow(s string, t string) string {
	targetHash := make(map[rune]int)
	for _, v := range t {
		targetHash[v]++
	}
	subHash := make(map[rune]int)
	res := ""
	l, r := 0, 0

	var isHashValid func() bool
	isHashValid = func() bool {
		for tKey, tVal := range targetHash {
			subVal := subHash[tKey]
			if subVal < tVal {
				return false
			}
		}
		return true
	}

	for r < len(s) {
		subHash[rune(s[r])]++
		for isHashValid() {
			if r-l+1 < len(res) || res == "" {
				res = s[l : r+1]
			}
			subHash[rune(s[l])]--
			l++
		}
		r++
	}
	return res
}
