package _8_validate_binary_search_tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	Node  *TreeNode
	Upper int
	Lower int
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*Queue{&Queue{
		Node:  root,
		Upper: math.MaxInt64,
		Lower: math.MinInt64,
	}}

	for len(queue) > 0 {
		for _, q := range queue {
			queue = queue[1:]
			if q.Node.Left != nil {
				queue = append(queue, &Queue{
					Node:  q.Node.Left,
					Upper: q.Node.Val,
					Lower: q.Lower,
				})
				if q.Node.Val <= q.Node.Left.Val || q.Lower >= q.Node.Left.Val {
					return false
				}
			}
			if q.Node.Right != nil {
				queue = append(queue, &Queue{
					Node:  q.Node.Right,
					Upper: q.Upper,
					Lower: q.Node.Val,
				})
				if q.Node.Val >= q.Node.Right.Val || q.Upper <= q.Node.Right.Val {
					return false
				}
			}
		}
	}
	return true
}
