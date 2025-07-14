package _24_binary_tree_maximum_path

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := -1001

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		if l < 0 {
			l = 0
		}
		r := dfs(node.Right)
		if r < 0 {
			r = 0
		}
		total := l + r + node.Val
		if total > res {
			res = total
		}
		return node.Val + max(l, r)
	}
	dfs(root)

	return res
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}
