package main

import (
	"fmt"
	"testing"
)

func Test_maximumUnits(t *testing.T) {
	in := [][]int{{2, 1}, {4, 4}, {3, 1}, {4, 1}, {2, 4}, {3, 4}, {1, 3}, {4, 3}, {5, 3}, {5, 3}}
	fmt.Println(maximumUnits(in, 13))
}
