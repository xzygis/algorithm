package _139_word_break

func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for _, word := range wordDict {
		m[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true //空字符串可以被拆分
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			// 划分为[0:j]和[j:i]这两段
			if dp[j] && m[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
