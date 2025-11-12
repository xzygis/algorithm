package _121_best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	profit := 0
	lowest := prices[0]
	for _, price := range prices {
		profit = max(profit, price-lowest)
		lowest = min(lowest, price)
	}

	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
