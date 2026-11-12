package _32_reconstruct_itinerary

import "sort"

func findItinerary(tickets [][]string) []string {
	// descendant lexical order
	// so that we can make stack
	sort.Slice(tickets, func(i, j int) bool {
		return tickets[i][1] > tickets[j][1]
	})

	graph := make(map[string][]string)

	for _, ticket := range tickets {
		graph[ticket[0]] = append(graph[ticket[0]], ticket[1])
	}

	res := []string{}

	var dfs func(src string)

	dfs = func(src string) {
		for len(graph[src]) > 0 {
			// pop from stack
			n := len(graph[src])
			dest := graph[src][n-1]
			graph[src] = graph[src][:n-1]
			dfs(dest)
		}
		res = append(res, src)
	}
	dfs("JFK")

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return res
}
