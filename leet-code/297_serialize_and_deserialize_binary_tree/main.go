package _97_serialize_and_deserialize_binary_tree

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	q := []*TreeNode{root}
	res := []string{}
	for len(q) > 0 {
		pop := q[0]
		q = q[1:]
		if pop == nil {
			res = append(res, "n")
			continue
		}
		res = append(res, fmt.Sprintf("%d", pop.Val))
		q = append(q, pop.Left)
		q = append(q, pop.Right)
	}
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	in := strings.Split(data, ",")
	root := &TreeNode{Val: toInt(in[0])}
	q := []*TreeNode{root}
	i := 1
	for len(q) > 0 {
		pop := q[0]
		q = q[1:]
		if in[i] != "n" {
			pop.Left = &TreeNode{Val: toInt(in[i])}
			q = append(q, pop.Left)
		}
		i++
		if in[i] != "n" {
			pop.Right = &TreeNode{Val: toInt(in[i])}
			q = append(q, pop.Right)
		}
		i++
	}
	return root
}

func toInt(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}
