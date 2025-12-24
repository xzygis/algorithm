package _72_edit_distance

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	if m*n == 0 {
		return m + n
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			v1 := dp[i-1][j] + 1
			v2 := dp[i][j-1] + 1
			v3 := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				v3 += 1
			}

			dp[i][j] = min(v1, v2, v3)
		}
	}

	return dp[m][n]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}
