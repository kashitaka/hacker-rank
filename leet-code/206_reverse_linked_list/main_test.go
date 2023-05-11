package _06_reverse_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	list13 := &ListNode{Val: 4, Next: nil}
	list12 := &ListNode{Val: 2, Next: list13}
	list11 := &ListNode{Val: 1, Next: list12}
	assert.Equal(t, list13, reverseList(list11))
}
