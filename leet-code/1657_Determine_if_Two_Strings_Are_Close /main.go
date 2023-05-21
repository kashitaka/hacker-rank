package _657_Determine_if_Two_Strings_Are_Close_

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	charFreq1, charFreq2 := make(map[rune]int), make(map[rune]int)
	// O(N)
	for _, v := range word1 {
		charFreq1[v]++
	}
	// O(M)
	for _, v := range word2 {
		charFreq2[v]++
	}
	// O(N)
	for k := range charFreq1 {
		_, ok := charFreq2[k]
		if !ok {
			return false
		}
	}


	numFreq1, numFreq2 := make(map[int]int), make(map[int]int)
	// O(N)
	for _, v := range charFreq1 {
		numFreq1[v]++
	}
	// O(M)
	for _, v := range charFreq2 {
		numFreq2[v]++
	}
	// O(N)
	for k, v := range numFreq1 {
		v2 := numFreq2[k]
		if v2 != v {
			return false
		}
	}
	// O(4N + 2M) = O(N+M)
	return true
}