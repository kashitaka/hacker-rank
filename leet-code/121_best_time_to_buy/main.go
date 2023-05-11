package _21_best_time_to_buy

func maxProfit(prices []int) int {
	maxPrft := 0
	min := 10000
	for _, v := range prices {
		if v < min {
			min = v
		} else {
			profit := v - min
			if profit > maxPrft {
				maxPrft = profit
			}
		}
	}
	return maxPrft
}
