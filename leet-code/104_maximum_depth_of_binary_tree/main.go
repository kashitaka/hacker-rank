package _04_maximum_depth_of_binary_tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root, 0)
}

func dfs(node *TreeNode, level int) int {
	if node.Left == nil && node.Right == nil{
		return level+1
	}
	resLeft, resRight := 0, 0
	if node.Left != nil {
		resLeft = dfs(node.Left, level+1)
	}
	if node.Right != nil {
		resRight = dfs(node.Right, level+1)
	}
	return max(resLeft, resRight)
}

func max(n, m int) int {
	if n >= m {
		return n
	}
	return m
}