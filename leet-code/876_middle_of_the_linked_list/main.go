package _76_middle_of_the_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	cur := head
	i := 0
	for head.Next != nil {
		if i%2 == 0 {
			cur = cur.Next
		}
		i++
		head = head.Next
	}
	return cur
}