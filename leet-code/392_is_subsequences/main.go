package _92_is_subsequences

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) > len(t) {
		return false
	}
	curS, curT := 0, 0
	for curS < len(s) {
		if curT >= len(t) {
			return false
		}
		charS := s[curS]
		for curT < len(t) {
			charT := t[curT]
			curT++
			if charT == charS {
				break
			}
			if curT >= len(t) {
				return false
			}
		}
		curS++
	}
	return true
}

func isSubsequence2(s string, t string) bool {
	curS, curT := 0, 0
	for curS < len(s) && curT < len(t) {
		if s[curS] == t[curT] {
			curS++
		}
		curT++
	}
	return curS == len(s)
}
