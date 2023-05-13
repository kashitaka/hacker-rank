package _89_n_ary_tree_preorder_traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_preorder(t *testing.T) {
	node5 := &Node{Val: 5, Children: nil}
	node6 := &Node{Val: 6, Children: nil}
	node3 := &Node{Val: 3, Children: []*Node{node5, node6}}
	node2 := &Node{Val: 2, Children: nil}
	node4 := &Node{Val: 4, Children: nil}
	node1 := &Node{Val: 1, Children: []*Node{node3, node2, node4}}
	assert.Equal(t, []int{1, 3, 5, 6, 2, 4}, preorder(node1))
}
