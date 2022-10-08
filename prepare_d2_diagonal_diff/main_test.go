package main

import (
	"fmt"
	"testing"
)

func Test_lonelyinteger(t *testing.T) {
	in := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	fmt.Println(diagonalDifference(in))
}
