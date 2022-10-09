package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	in := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	fmt.Println(numIslands(in))
}
