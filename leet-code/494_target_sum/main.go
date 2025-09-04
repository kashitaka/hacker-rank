package _94_target_sum

func findTargetSumWays(nums []int, target int) int {
	dp := make(map[int]int)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		newDP := make(map[int]int)
		num := nums[i]
		for k, v := range dp {
			newDP[k-num] += v
			newDP[k+num] += v
		}
		dp = newDP
	}
	return dp[target]
}
