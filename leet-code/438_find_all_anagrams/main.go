package _38_find_all_anagrams

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	size := len(p)
	pMap := make(map[int32]int)
	for _, v := range p{
		pMap[v]++
	}
	sMap := make(map[int32]int)
	for _, v := range s[0:size] {
		sMap[v]++
	}
	var res []int
	if equal(pMap, sMap) {
		res = append(res, 0)
	}
	for i := 1; i <= len(s) - size; i++ {
		sMap[int32(s[i-1])]--
		sMap[int32(s[i+size-1])]++
		if equal(pMap, sMap) {
			res = append(res, i)
		}
	}
	return res
}

func equal(pMap, sMap map[int32]int) bool {
	for k, v := range pMap {
		if sMap[k] != v {
			return false
		}
	}
	return true
}