package main

import "fmt"

var digitsLetter = map[int32][]string{
	50: {"a", "b", "c"},
	51: {"d", "e", "f"},
	52: {"g", "h", "i"},
	53: {"j", "k", "l"},
	54: {"m", "n", "o"},
	55: {"p", "q", "r", "s"},
	56: {"t", "u", "v"},
	57: {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res := []string{""}
	for _, v := range digits {
		lettersToAdd := digitsLetter[v]
		var newRes []string
		for _, v := range res {
			for _, add := range lettersToAdd {
				newRes = append(newRes, fmt.Sprint(v)+add)
				res = newRes
			}
		}
	}
	return res
}
