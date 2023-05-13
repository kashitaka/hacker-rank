package _89_n_ary_tree_preorder_traversal

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	return dfs(root)
}

func dfs(node *Node) []int {
	var res = []int{node.Val}
	for _, v := range node.Children {
		res = append(res, dfs(v)...)
	}
	return res
}
