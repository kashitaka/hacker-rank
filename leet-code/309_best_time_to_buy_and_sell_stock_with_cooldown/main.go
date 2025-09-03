package _09_best_time_to_buy_and_sell_stock_with_cooldown

func maxProfit(prices []int) int {
	// dp[i] is max profit of subArray prices[i:]
	// dp[i][0] is the case where we BUY (or cooldown) in ith day
	// dp[i][1] is the case where we SELL (or cooldown) in ith day
	dp := make([][2]int, len(prices))
	for i, _ := range dp {
		dp[i] = [2]int{0, 0} // idx 0: buying / idx 1: selling
	}

	// dfs returns max profit of subArray prices[i:]
	// in the case where given argument: buying: 0 = buying/1 ==selling
	var dfs func(i int, buying int) int
	dfs = func(i int, buying int) int {
		if i >= len(prices) {
			return 0
		}
		if dp[i][buying] != 0 {
			return dp[i][buying]
		}
		if buying == 0 {
			// we can buy or cooldown
			cooldown := dfs(i+1, buying)
			buy := dfs(i+1, 1) - prices[i]
			dp[i][buying] = max(cooldown, buy)
		} else {
			// we can sell or cooldown
			cooldown := dfs(i+1, buying)
			sell := dfs(i+2, 0) + prices[i]
			dp[i][buying] = max(cooldown, sell)
		}
		return dp[i][buying]
	}
	return dfs(0, 0)
}
