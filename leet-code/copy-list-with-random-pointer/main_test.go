package main

import "testing"

func Test(t *testing.T) {
	n3 := Node{
		Val:  3,
		Next: nil,
	}
	n2 := Node{
		Val:  2,
		Next: &n3,
	}
	n1 := Node{
		Val:  1,
		Next: &n2,
	}
	n2.Random = &n1
	copyRandomList(&n1)
}
