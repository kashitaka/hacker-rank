package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	in := [][]int{{0, 30}, {5, 10}, {15, 20}}
	fmt.Println(minMeetingRooms(in))
}
