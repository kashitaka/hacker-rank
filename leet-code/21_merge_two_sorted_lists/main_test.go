package _1_merge_two_sorted_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	list13 := &ListNode{Val: 4, Next: nil}
	list12 := &ListNode{Val: 2, Next: list13}
	list11 := &ListNode{Val: 1, Next: list12}
	list23 := &ListNode{Val: 4, Next: nil}
	list22 := &ListNode{Val: 3, Next: list23}
	list21 := &ListNode{Val: 1, Next: list22}
	assert.Equal(t, list11, mergeTwoLists(list11, list21))
}
