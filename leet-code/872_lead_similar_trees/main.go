package _72_lead_similar_trees

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var leaf1, leaf2 []int
	dfs(root1, &leaf1)
	dfs(root2, &leaf2)
	if len(leaf1) != len(leaf2) {
		return false
	}
	for i, v := range leaf1 {
		if leaf2[i] != v {
			return false
		}
	}
	return true
}

func dfs(node *TreeNode, leaves *[]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		*leaves = append(*leaves, node.Val)
		return
	}
	if node.Left != nil {
		dfs(node.Left, leaves)
	}
	if node.Right != nil {
		dfs(node.Right, leaves)
	}
}


