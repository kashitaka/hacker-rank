package _1_merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	previous := preHead

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			previous.Next = list1
			list1 = list1.Next
		} else {
			previous.Next = list2
			list2 = list2.Next
		}
		previous = previous.Next
	}

	// 残ったlistに接続して終わり
	if list1 == nil {
		previous.Next = list2
	} else {
		previous.Next = list1
	}

	return preHead.Next
}
