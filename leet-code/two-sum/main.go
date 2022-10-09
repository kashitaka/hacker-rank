package main

func twoSum(nums []int, target int) []int {
	exist := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		v1 := nums[i]
		i2, ok := exist[target-v1]
		if ok && i2 != i {
			return []int{i, i2}
		}
		exist[v1] = i
	}
	panic("impossible")
}
