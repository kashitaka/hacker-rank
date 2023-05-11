package _42_linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	checkMap := make(map[*ListNode]struct{})
	for head != nil {
		_, ok := checkMap[head]
		if ok {
			return head
		}
		checkMap[head] = struct{}{}
		head = head.Next
	}
	return nil
}
