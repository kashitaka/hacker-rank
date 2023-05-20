package _768_merge_strings_alternatly

func mergeAlternately(word1 string, word2 string) string {
	res := []byte{}
	cur := 0
	for cur < len(word1) && cur < len(word2) {
		res = append(res, word1[cur])
		res = append(res, word2[cur])
		cur++
	}
	resStr := string(res)
	if len(word1) > len(word2) {
		return resStr + word1[cur:]
	}
	return resStr + word2[cur:]
}