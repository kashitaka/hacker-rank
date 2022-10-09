package main

import "fmt"

var order []int
var circle bool

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
	}
	seen := make([]bool, len(graph))
	for i := range seen {
		seen[i] = false
	}
	circle = false
	order = []int{}
	fmt.Println(order)
	for i := range graph {
		if seen[i] {
			continue
		}
		rec(graph, seen, i, i)
	}
	fmt.Println(order)
	if circle {
		return []int{}
	}
	for i := 0; i < len(order)/2; i++ {
		order[i], order[len(order)-1-i] = order[len(order)-1-i], order[i]
	}
	fmt.Println(order)
	return order
}

func rec(graph [][]int, seen []bool, start, v int) {
	seen[v] = true
	nextVs := graph[v]
	for _, nv := range nextVs {
		if seen[nv] {
			circle = true
			continue
		}
		rec(graph, seen, start, nv)
	}
	order = append(order, v)
}
