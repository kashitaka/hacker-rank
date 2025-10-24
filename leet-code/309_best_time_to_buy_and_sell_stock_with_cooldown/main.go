package _09_best_time_to_buy_and_sell_stock_with_cooldown

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	hold, rest, sell := -prices[0], 0, -10000

	for i := 1; i < len(prices); i++ {
		prevHold, prevRest, prevSell := hold, rest, sell
		hold = max(prevHold, prevRest-prices[i])
		rest = max(prevRest, prevSell)
		sell = prevHold + prices[i]
	}
	return max(rest, sell)
}
