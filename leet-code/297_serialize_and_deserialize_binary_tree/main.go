package _97_serialize_and_deserialize_binary_tree

import (
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
func (c *Codec) serializeBFS(root *TreeNode) string {
	res := make([]string, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			res = append(res, "n")
		}
		if node != nil {
			res = append(res, strconv.Itoa(node.Val))
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (c *Codec) deserializeBFS(data string) *TreeNode {
	list := strings.Split(data, ",")
	if list[0] == "n" {
		return nil
	}
	v, err := strconv.Atoi(list[0])
	if err != nil {
		panic(err)
	}
	root := &TreeNode{
		Val: v,
	}
	queue := []*TreeNode{root}
	idx := 1

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		strL := list[idx]
		var nodeL *TreeNode
		if strL != "n" {
			v, e := strconv.Atoi(list[idx])
			if e != nil {
				panic(e)
			}

			nodeL = &TreeNode{Val: v}
			queue = append(queue, nodeL)
		}
		node.Left = nodeL
		idx++

		strR := list[idx]
		var nodeR *TreeNode
		if strR != "n" {
			v, e := strconv.Atoi(list[idx])
			if e != nil {
				panic(e)
			}
			nodeR = &TreeNode{Val: v}
			queue = append(queue, nodeR)
		}
		node.Right = nodeR
		idx++
	}

	return root
}

func (c *Codec) serializeDFS(root *TreeNode) string {
	res := make([]string, 0)

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			res = append(res, "n")
			return
		}
		v := strconv.Itoa(node.Val)
		res = append(res, v)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (c *Codec) deserializeDFS(data string) *TreeNode {
	list := strings.Split(data, ",")
	idx := 0
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		strV := list[idx]
		idx++
		if strV == "n" {
			return nil
		}
		intV, _ := strconv.Atoi(strV)
		node := &TreeNode{Val: intV}
		node.Left = dfs()
		node.Right = dfs()
		return node
	}
	return dfs()
}
