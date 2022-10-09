package main

import "testing"

func Test(t *testing.T) {
	findOrder(4, [][]int{{2, 0}, {1, 0}, {3, 1}, {3, 2}, {1, 3}})
}
