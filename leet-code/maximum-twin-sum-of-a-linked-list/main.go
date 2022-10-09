package main

func pairSum(head *ListNode) int {
	var valList []int
	cur := head
	for cur != nil {
		valList = append(valList, cur.Val)
		cur = cur.Next
	}
	mid := len(valList) / 2
	max := 0
	for i := 0; i < mid; i++ {
		sum := valList[i] + valList[len(valList)-1-i]
		if sum > max {
			max = sum
		}
	}
	return max
}
