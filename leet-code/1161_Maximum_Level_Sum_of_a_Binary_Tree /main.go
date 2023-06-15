package _161_Maximum_Level_Sum_of_a_Binary_Tree_

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	queue := make([]*TreeNode, 1)
	queue[0] = root

	var sumSlice []int
	level := 0
	for len(queue) > 0 {
		sumSlice = append(sumSlice, 0)
		for _, v := range queue {
			queue = queue[1:]
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
			sumSlice[level] += v.Val
		}
		level++
	}

	max := -100000
	maxIndex := 0

	for i, v := range sumSlice {
		if v > max {
			max = v
			maxIndex = i
		}
	}

	return maxIndex + 1
}
