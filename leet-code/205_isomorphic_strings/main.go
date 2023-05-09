package _05_isomorphic_strings

func isIsomorphic(s string, t string) bool {
	mapStoT := make(map[uint8]uint8)
	mapTtoS := make(map[uint8]uint8)
	for i, v := range s {
		char, ok1 := mapStoT[uint8(v)]
		if ok1 && t[i] != char {
			return false
		}
		reverse, ok2 := mapTtoS[t[i]]
		if ok2 && reverse != uint8(v) {
			return false
		}
		if !ok1 {
			mapStoT[uint8(v)] = t[i]
			mapTtoS[t[i]] = uint8(v)
		}
	}
	return true
}
