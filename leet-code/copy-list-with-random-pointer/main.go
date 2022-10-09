package main

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	oldNewMap := make(map[*Node]*Node)
	newHead := openNode(head, oldNewMap)
	for k, v := range oldNewMap {
		oldRandom := k.Random
		newRandom := oldNewMap[oldRandom]
		v.Random = newRandom
	}
	return newHead
}

func openNode(node *Node, oldNewMap map[*Node]*Node) *Node {
	if node.Next == nil {
		n := &Node{
			Val:    node.Val,
			Next:   nil,
			Random: nil,
		}
		oldNewMap[node] = n
		return n
	}
	newNext := openNode(node.Next, oldNewMap)
	newNode := Node{
		Val:    node.Val,
		Next:   newNext,
		Random: nil,
	}
	oldNewMap[node] = &newNode
	return &newNode
}
