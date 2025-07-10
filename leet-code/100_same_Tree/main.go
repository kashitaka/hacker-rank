package main

import leetcode "hacker-rank/leet-code"

func isSameTree(p *leetcode.TreeNode, q *leetcode.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil && p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}
