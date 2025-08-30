package _22_reconstruct_itinerary

import "sort"

func findItinerary(tickets [][]string) []string {
	sort.Slice(tickets, func(i, j int) bool {
		return tickets[i][1] > tickets[j][1]
	})
	adj := make(map[string][]string)
	for _, v := range tickets {
		adj[v[0]] = append(adj[v[0]], v[1])
	}
	res := []string{}

	var dfs func(src string)
	dfs = func(src string) {
		for len(adj[src]) > 0 {
			n := len(adj[src])
			pop := adj[src][n-1]
			adj[src] = adj[src][:n-1]
			dfs(pop)
		}
		res = append(res, src)
	}
	dfs("JFK")

	// reverse order in res
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return res
}
