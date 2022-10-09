package main

import "fmt"

type Graph struct {
	g   [][]int
	res []int
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	if k == 0 {
		return []int{target.Val}
	}

	// convert to graph
	graph := &Graph{}
	g := make([][]int, 500)
	var res []int
	graph.g = g
	graph.res = res
	open(root.Val, root.Left, graph)
	open(root.Val, root.Right, graph)
	fmt.Println(graph.g)

	// 探索
	seen := make([]bool, 500)

	dfs(target.Val, 0, k, graph, seen)
	return graph.res
}

func open(parentVal int, node *TreeNode, graph *Graph) {
	if node == nil {
		return
	}
	thisVal := node.Val
	graph.g[parentVal] = append(graph.g[parentVal], thisVal)
	graph.g[thisVal] = append(graph.g[thisVal], parentVal)

	open(thisVal, node.Left, graph)
	open(thisVal, node.Right, graph)
}

func dfs(i int, depth int, k int, graph *Graph, seen []bool) {
	seen[i] = true
	nextVs := graph.g[i]
	fmt.Println(nextVs, depth)
	depth++
	for _, v := range nextVs {
		if seen[v] {
			continue
		}
		if depth == k {
			graph.res = append(graph.res, v)
			continue
		}
		dfs(v, depth, k, graph, seen)
	}
}
