package _38_copy_list_with_random_ptr

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	hash := make(map[*Node]*Node)
	hash[nil] = nil
	cur := head
	for cur != nil {
		hash[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		hash[cur].Next = hash[cur.Next]
		hash[cur].Random = hash[cur.Random]
		cur = cur.Next
	}

	return hash[head]
}
