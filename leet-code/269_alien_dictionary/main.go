package _69_alien_dictionary

func alienOrder(words []string) string {
	adj := make(map[rune][]rune)
	for _, word := range words {
		for _, c := range word {
			adj[c] = []rune{}
		}
	}

	// make adj map
	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		minLen := min(len(w1), len(w2))
		if w1[:minLen] == w2[:minLen] && len(w1) > len(w2) {
			// same prefix but w1 is longer than w2
			return ""
		}
		for j := 0; j < minLen; j++ {
			if w1[j] != w2[j] {
				adj[rune(w1[j])] = append(adj[rune(w1[j])], rune(w2[j]))
				break
			}
		}
	}

	res := []rune{}
	visited := make(map[rune]bool) // null: not visited, false: visited, true: in dfs path

	var dfs func(rune) bool
	// return true if there is circular dependency, meaning invalid lixicographical order
	dfs = func(char rune) bool {
		if isVisit, ok := visited[char]; ok {
			// alrady in the path
			return isVisit
		}
		visited[char] = true

		for _, neighChar := range adj[char] {
			if dfs(neighChar) {
				return true
			}
		}
		visited[char] = false
		res = append(res, char)
		return false
	}

	for key, _ := range adj {
		if dfs(key) {
			return ""
		}
	}

	result := []byte{}
	for i := len(res) - 1; i >= 0; i-- {
		result = append(result, byte(res[i]))
	}
	return string(result)
}
