package _16_partition_equal_subset_sum

func canPartition(nums []int) bool {
	total := 0
	for _, v := range nums {
		total += v
	}
	if total%2 == 1 {
		return false
	}
	t := total / 2
	sums := make(map[int]struct{})
	sums[0] = struct{}{}
	for i := 0; i < len(nums); i++ {
		newKeys := []int{}
		for key, _ := range sums {
			if key+nums[i] == t {
				return true
			}
			newKeys = append(newKeys, key+nums[i])
		}
		for _, v := range newKeys {
			sums[v] = struct{}{}
		}
	}
	return false
}
