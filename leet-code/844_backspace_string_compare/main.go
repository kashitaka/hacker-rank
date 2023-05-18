package _44_backspace_string_compare

func backspaceCompare(s string, t string) bool {
	var sStack []uint8
	var tStack []uint8

	for i:=0; i<len(s); i++ {
		if s[i] == 35 {
			// #
			if len(sStack) > 0 {
				sStack = sStack[0:len(sStack)-1]
			}
			continue
		}
		sStack = append(sStack, s[i])
	}
	for i:=0; i<len(t); i++ {
		if t[i] == 35 {
			// #
			if len(tStack) > 0 {
				tStack = tStack[0:len(tStack)-1]
			}
			continue
		}
		tStack = append(tStack, t[i])
	}
	if len(sStack) != len(tStack) {
		return false
	}
	for i, v := range sStack {
		if tStack[i] != v {
			return false
		}
	}
	return true
}