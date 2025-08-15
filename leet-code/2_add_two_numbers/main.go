package __add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n1 + n2)
// Space: O(1)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	dummy := &ListNode{}
	prev := dummy
	for l1 != nil || l2 != nil || carry != 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		s := v1 + v2 + carry
		if s >= 10 {
			carry = 1
			s -= 10
		} else {
			carry = 0
		}
		node := &ListNode{Val: s}
		prev.Next = node
		prev = node
	}
	return dummy.Next
}
