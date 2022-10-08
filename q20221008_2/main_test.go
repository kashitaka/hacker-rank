package main

import (
	"fmt"
	"testing"
)

func Test_count(t *testing.T) {
	s := "**|****|*|**|*"
	res := count(s, 4, 13)
	fmt.Println(res)
}

func Test_searchMem(t *testing.T) {
	s := "**|****|*|**|*"
	mem := make(map[int32][]int32)
	createMem(s, mem)
	res := count2(4-1, 13-1, mem)
	fmt.Println(res)
}
