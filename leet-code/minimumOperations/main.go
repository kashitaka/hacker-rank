package main

func minimumOperations(nums []int) int {
	exist := make(map[int]struct{})
	for _, v := range nums {
		if v > 0 {
			exist[v] = struct{}{}
		}
	}
	return len(exist)
}
