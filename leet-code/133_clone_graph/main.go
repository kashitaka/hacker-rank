package _33_clone_graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	seen := make(map[int]*Node)
	root := &Node{
		Val: node.Val,
	}
	seen[root.Val] = root

	var dfs func(n *Node)
	dfs = func(n *Node) {
		copy := seen[n.Val]
		if copy != nil && len(copy.Neighbors) > 0 {
			// already visited
			return
		}
		if copy == nil {
			copy = &Node{
				Val: n.Val,
			}
			seen[copy.Val] = copy
		}
		// make copy neighbors
		neighbors := make([]*Node, len(n.Neighbors))
		for i, v := range n.Neighbors {
			neigbor := seen[v.Val]
			if neigbor == nil {
				neigbor = &Node{
					Val: v.Val,
				}
				seen[v.Val] = neigbor
			}
			neighbors[i] = neigbor
		}
		copy.Neighbors = neighbors
		for _, v := range n.Neighbors {
			dfs(v)
		}
	}
	dfs(node)
	return root
}

// NeetCode 参考に。
// 頭いい
func cloneGraphImproved(node *Node) *Node {
	if node == nil {
		return nil
	}
	nodeHash := make(map[int]*Node) // map of originalVal vs Node. it works because vals are unique.

	var dfs func(*Node) *Node
	dfs = func(orig *Node) *Node {
		if node, ok := nodeHash[orig.Val]; ok {
			return node
		}

		newNode := &Node{

			Val: orig.Val,
		}
		nodeHash[newNode.Val] = newNode
		neighbors := make([]*Node, len(orig.Neighbors))
		for i, v := range orig.Neighbors {
			neighbors[i] = dfs(v)
		}
		newNode.Neighbors = neighbors
		return newNode
	}
	return dfs(node)
}
