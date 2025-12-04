package _279_perfect_squares

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minPart := math.MaxInt
		for j := 1; j*j <= i; j++ {
			minPart = min(minPart, dp[i-j*j])
		}
		dp[i] = minPart + 1
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
