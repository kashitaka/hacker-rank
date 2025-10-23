package _16_partition_equal_subset_sum

func canPartition(nums []int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}
	if total%2 != 0 {
		return false
	}

	target := total / 2

	// dp[i] は nums を使って整数iを作れるかの真理値
	dp := make([]bool, target+1)
	// dp[0] は常に true。なぜなら0はnumsを使って(何も使わずに)作れるため
	dp[0] = true
	for _, num := range nums {
		for i := target; i >= num; i-- {
			if dp[i-num] {
				dp[i] = true
			}
		}
	}
	return dp[target]
}
