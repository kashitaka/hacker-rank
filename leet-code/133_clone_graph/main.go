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
func cloneGraphImprove(node *Node) *Node {
	if node == nil {
		return nil
	}
	oldNew := make(map[*Node]*Node)

	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		cn := oldNew[node]
		if cn != nil {
			return cn
		}
		cn = &Node{
			Val: node.Val,
		}
		oldNew[node] = cn
		copyNeigbors := make([]*Node, len(node.Neighbors))
		for idx, neighbor := range node.Neighbors {
			copyNeigbors[idx] = dfs(neighbor)
		}
		cn.Neighbors = copyNeigbors
		return cn
	}

	return dfs(node)
}
