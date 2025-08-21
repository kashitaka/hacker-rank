package _7_letter_combinations_of_a_phone_number

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	numChar := map[byte][]string{
		'2': {"a", "b", "c"}, '3': {"d", "e", "f"}, '4': {"g", "h", "i"},
		'5': {"j", "k", "l"}, '6': {"m", "n", "o"}, '7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"}, '9': {"w", "x", "y", "z"},
	}

	res := []string{}
	var dfs func(idx int, str string)
	dfs = func(idx int, str string) {
		if idx == len(digits) {
			res = append(res, str)
			return
		}
		candidates := numChar[digits[idx]]
		for _, char := range candidates {
			newStr := str + char
			dfs(idx+1, newStr)
		}
	}
	dfs(0, "")
	return res
}
