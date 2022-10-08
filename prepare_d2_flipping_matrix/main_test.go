package main

import (
	"fmt"
	"testing"
)

func Test_lonelyinteger(t *testing.T) {
	in := []int32{1, 2, 3, 4, 3, 2, 1}
	fmt.Println(lonelyinteger(in))
}
