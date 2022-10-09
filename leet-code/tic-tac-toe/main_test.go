package main

import "testing"

func Test(t *testing.T) {
	tic := Constructor(3)
	tic.Move(0, 0, 1)
	tic.Move(0, 2, 2)
	tic.Move(2, 2, 1)
	tic.Move(1, 1, 2)
}
