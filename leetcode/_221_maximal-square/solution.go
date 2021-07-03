package _221_maximal_square

func maximalSquare(matrix [][]byte) int {
	rows := len(matrix)
	cols := len(matrix[0])

	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}

	maxLen := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(min(dp[i][j-1], dp[i-1][j-1]), dp[i-1][j]) + 1
				}
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
				}
			}
		}
	}

	return maxLen * maxLen
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}