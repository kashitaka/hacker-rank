package _35_lowest_common_ancestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

func dfs1(node, p, q *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if p.Val <= node.Val && node.Val <= q.Val {
		return node
	}
	res := dfs1(node.Left, p, q)
	if res != nil {
		return res
	}
	res = dfs1(node.Right, p, q)
	if res != nil {
		return res
	}
	return nil
}
