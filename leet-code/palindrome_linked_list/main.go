package main

func isPalindrome(head *ListNode) bool {
	var list []int32
	list = append(list, int32(head.Val))
	cur := head.Next
	for cur != nil {
		node := cur
		list = append(list, int32(node.Val))
		cur = node.Next
	}
	for i := 0; i < len(list)/2; i++ {
		if list[i] != list[len(list)-1-i] {
			return false
		}
	}
	return true

}
