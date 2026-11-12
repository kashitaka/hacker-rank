package _69_alien_dictionary

// Let n be the length of given words
// Suppose each word length is at most l.
// Overall running time is
// O(nl + nl) = O(nl)
func alienOrder(words []string) string {
	graph := make(map[rune]map[rune]struct{})

	// O(nl) for initializing the graph
	for _, word := range words {
		for _, char := range word {
			if _, ok := graph[char]; !ok {
				graph[char] = make(map[rune]struct{})
			}
		}
	}

	// make (n-1) pairs
	// for each word, compare at most l characters
	// so, O(nl)
	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		if len(w1) > len(w2) && w2 == w1[:len(w2)] {
			return ""
		}

		for k := 0; k < len(w1) && k < len(w2); k++ {
			if w1[k] != w2[k] {
				graph[rune(w1[k])][rune(w2[k])] = struct{}{}
				break
			}
		}
	}

	res := []rune{}

	// Topological Sort
	// For a given graph G=(V, E), check all V and E,
	// so the running time is O(|V|+ |E|).
	// Since the nodes are lowercase English letters, |V| is at most 26.
	// For each V, there are at most 26-1 edges, so the number of edges
	// is at most 25 * 26.
	// Therefore, O(|V|+|E|) is somewhat huge but still constant time O(1)
	visit, visiting := make(map[rune]bool), make(map[rune]bool)
	var dfs func(rune) bool
	dfs = func(src rune) bool {
		if visit[src] {
			return true
		}

		if visiting[src] {
			return false
		}

		visiting[src] = true

		// process
		for dest, _ := range graph[src] {
			if !dfs(dest) {
				return false
			}
		}

		// tear down
		visiting[src] = false
		visit[src] = true
		res = append(res, src)
		return true
	}

	for src, _ := range graph {
		if !dfs(src) {
			return ""
		}
	}

	// O(26/2) = O(1)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return string(res)
}
