package _10_balanced_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	res := true
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Right)
		r := dfs(node.Left)
		if l-r > 1 || -1 > l-r {
			res = false
		}
		return 1 + max(l, r)
	}
	dfs(root)
	return res
}
