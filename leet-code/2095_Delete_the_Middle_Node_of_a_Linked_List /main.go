package _095_Delete_the_Middle_Node_of_a_Linked_List_

type ListNode struct {
	Val int
	Next *ListNode
}

// deleteMiddle returns.
// O(N)だが遅い。LinkedListの中間ノードは fast, slow ポインタを使うと
// hash 使わなくても簡単に特定できる
func deleteMiddle(head *ListNode) *ListNode {
	indexNodeMap := make(map[int]*ListNode)
	i := 0
	// O(N)
	for head != nil {
		indexNodeMap[i] = head
		head = head.Next
		i++
	}
	if i == 1 {
		return nil
	}
	middleNode := indexNodeMap[i/2]
	prevNode := indexNodeMap[i/2-1]
	if middleNode.Next == nil {
		prevNode.Next = nil
	} else {
		prevNode.Next = middleNode.Next
	}
	return indexNodeMap[0]
}