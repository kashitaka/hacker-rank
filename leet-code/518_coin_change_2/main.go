package _18_coin_change_2

func change(amount int, coins []int) int {
	dp := make([][]int, len(coins))
	for i, _ := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}

	for i := len(coins) - 1; i >= 0; i-- {
		coin := coins[i]
		for j := 1; j <= amount; j++ {
			if j-coin >= 0 {
				dp[i][j] += dp[i][j-coin]
			}
			if i < len(coins)-1 {
				dp[i][j] += dp[i+1][j]
			}
		}
	}
	return dp[0][amount]
}
