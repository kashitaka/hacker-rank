package main

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	node2 := ListNode{
		Val:  2,
		Next: nil,
	}
	node1 := ListNode{
		Val:  1,
		Next: &node2,
	}
	fmt.Println(isPalindrome(&node1))
}
