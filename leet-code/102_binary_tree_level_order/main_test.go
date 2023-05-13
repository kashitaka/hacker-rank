package _02_binary_tree_level_order

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_bfs(t *testing.T) {
	n15 := &TreeNode{Val: 15, Left: nil, Right: nil}
	n7 := &TreeNode{Val: 7, Left: nil, Right: nil}
	n20 := &TreeNode{Val: 20, Left: n15, Right: n7}
	n9 := &TreeNode{Val: 9, Left: nil, Right: nil}
	n3 := &TreeNode{Val: 3, Left: n9, Right: n20}
	assert.Equal(t, [][]int{{3}, {9, 20}, {15, 7}}, levelOrder(n3))
}
