package _24_longest_repeating_character

func characterReplacement(s string, k int) int {
	start, end := 0,k
	max := k
	windowSize := k
	freqMap := make(map[int32]int)
	for _, v := range s[0:k] {
		freqMap[v]++
	}
	for end < len(s) {
		end++
		for _, v := range s[end-1:end] {
			freqMap[v]++
		}
		subStr := s[start: end]
		if check(subStr, k, freqMap) {
			max = len(subStr)
			windowSize++
			continue
		}
		for _, v := range s[start:start+1] {
			freqMap[v]--
		}
		start++
	}
	return max
}

func check(subStr string, k int, freqMap map[int32]int) bool {
	max := 0
	for _, v := range freqMap {
		if v > max {
			max = v
		}
	}
	if len(subStr) - max <= k {
		return true
	}
	return false
}