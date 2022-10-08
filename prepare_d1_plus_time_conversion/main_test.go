package main

import (
	"fmt"
	"testing"
)

func Test_timeConversion(t *testing.T) {
	s := "02:01:00AM"
	fmt.Println(timeConversion(s))
}
