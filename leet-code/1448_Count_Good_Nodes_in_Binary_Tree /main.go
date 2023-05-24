package _448_Count_Good_Nodes_in_Binary_Tree_

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(node *TreeNode, max int) int {
	cnt := 0
	if node.Val >= max {
		cnt++
		max = node.Val
	}
	leftCnt, rightCnt := 0, 0
	if node.Left != nil {
		leftCnt = dfs(node.Left, max)
	}
	if node.Right != nil {
		rightCnt = dfs(node.Right, max)
	}
	return cnt + leftCnt + rightCnt
}
