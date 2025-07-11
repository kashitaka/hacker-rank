package _05_constaruct_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NeetCode の最適化。考え方はわかるが実装がむずすぎ
func buildTree_3(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	hashmap := make(map[int]int, len(inorder))
	for i, v := range inorder {
		hashmap[v] = i
	}
	preIdx := 0

	var dfs func(left int, right int) *TreeNode
	dfs = func(left int, right int) *TreeNode {
		if left > right {
			return nil
		}

		root := &TreeNode{Val: preorder[preIdx]}
		mid := hashmap[root.Val]
		preIdx++

		root.Left = dfs(left, mid-1)
		root.Right = dfs(mid+1, right)
		return root
	}

	return dfs(0, len(inorder)-1)
}

// NeetCode
func buildTree_2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	mid := 0
	for i, v := range inorder {
		if v == root.Val {
			mid = i
			break
		}
	}
	root.Left = buildTree_2(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree_2(preorder[1+mid:], inorder[1+mid:])
	return root
}

// first submission
func buildTree(preorder []int, inorder []int) *TreeNode {
	inoderIdx := make(map[int]int, len(inorder))
	for i, v := range inorder {
		inoderIdx[v] = i
	}

	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}

	for i := 1; i < len(preorder); i++ {
		atatchNode := &TreeNode{
			Val:   preorder[i],
			Left:  nil,
			Right: nil,
		}
		curr := root
		attatchNodeIdx := inoderIdx[preorder[i]]
		for curr != nil {
			currIdx := inoderIdx[curr.Val]
			if currIdx < attatchNodeIdx {
				// attachNode should be in right side
				if curr.Right == nil {
					// atach right
					curr.Right = atatchNode
					break
				} else {
					curr = curr.Right
				}
			} else {
				// atatchNode should be in left side
				if curr.Left == nil {
					// atatch left
					curr.Left = atatchNode
					break
				} else {
					curr = curr.Left
				}
			}
		}
	}
	return root
}
