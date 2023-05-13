package _02_binary_tree_level_order

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	var res [][]int
	level := 0

	for len(queue) > 0 {
		res = append(res, []int{})
		// ある分は全部回す。このループ中にqueueが増えるが、for range回数は増えない
		for _, node := range queue {
			res[level] = append(res[level], node.Val)
			// dequeue
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		level++
	}
	return res
}
