package main

import treenode "hacker-rank/leet-code/tree_node"

func isSameTree(p *treenode.TreeNode, q *treenode.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil && p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}
