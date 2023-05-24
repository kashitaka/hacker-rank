package _99_Binary_Tree_Right_Side_View_

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	var res []int
	for len(queue) > 0 {
		lastItem := queue[len(queue)-1].Val
		res = append(res, lastItem)
		for _, v := range queue {
			queue = queue[1:]
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
	}
	return res
}