package main

import "sort"

//func groupAnagrams(strs []string) [][]string {
//	resMap := make(map[string][]string)
//	for _, v := range strs {
//		key := make([]int, 26)
//		for j := range key {
//			key[j] = 0
//		}
//		runes := []rune(v)
//		for _, v := range runes {
//			key[v-97]++
//		}
//		resMap[fmt.Sprint(key)] = append(resMap[fmt.Sprint(key)], v)
//	}
//	res := make([][]string, len(resMap))
//	i := 0
//	for _, v := range resMap {
//		res[i] = v
//		i++
//	}
//	return res
//}

func groupAnagrams(strs []string) [][]string {
	sortedStrs := make([]string, len(strs))
	for i, v := range strs {
		runes := []rune(v)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		sortedStrs[i] = string(runes)
	}
	resMap := make(map[string][]string)
	for i, v := range sortedStrs {
		resMap[v] = append(resMap[v], strs[i])
	}
	res := make([][]string, len(resMap))
	i := 0
	for _, v := range resMap {
		res[i] = v
		i++
	}

	return res
}
