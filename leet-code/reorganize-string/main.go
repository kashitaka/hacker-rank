package main

import (
	"sort"
)

func reorganizeString(s string) string {
	hashmap := make(map[int32]int)
	for _, v := range s {
		hashmap[v] += 1
	}
	keys := make([]int32, 0, len(hashmap))
	for k := range hashmap {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return hashmap[keys[i]] > hashmap[keys[j]]
	})
	maxLetter := keys[0]
	maxLetterCount := hashmap[keys[0]]
	if (len(s)%2 == 0 && maxLetterCount > len(s)/2) ||
		(len(s)%2 == 1 && maxLetterCount > len(s)/2+1) {
		return ""
	}

	delete(hashmap, maxLetter)
	keys = keys[1:]

	res := make([]rune, len(s))
	idx := 0
	for maxLetterCount > 0 {
		res[idx] = maxLetter
		idx += 2
		maxLetterCount--
	}
	for k := range hashmap {
		for hashmap[k] > 0 {
			if idx >= len(res) {
				idx = 1
			}
			res[idx] = k
			hashmap[k]--
			idx += 2
		}
	}
	return string(res)
}
