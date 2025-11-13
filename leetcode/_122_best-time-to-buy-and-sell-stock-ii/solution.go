package _122_best_time_to_buy_and_sell_stock_ii

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	length := len(prices)
	dp := make([][2]int, length)
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[length-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxProfitV1(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}
