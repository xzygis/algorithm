package longestpalindromicsubstring

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	maxLen := 1
	begin := 0
	for paLen := 2; paLen <= len(s); paLen++ {
		for l := 0; l+paLen-1 < len(s); l++ {
			r := l + paLen - 1
			if s[l] != s[r] {
				dp[l][r] = false
			} else if r-l < 3 {
				dp[l][r] = true
			} else {
				dp[l][r] = dp[l+1][r-1]
			}

			if paLen > maxLen && dp[l][r] {
				maxLen = paLen
				begin = l
			}
		}
	}

	return s[begin : begin+maxLen]
}

func longestPalindromeV1(s string) string {
	var ans string
	for i := 0; i < len(s); i++ {
		palindome := helper(s, i, i+1)
		if len(palindome) > len(ans) {
			ans = palindome
		}

		palindome = helper(s, i, i)
		if len(palindome) > len(ans) {
			ans = palindome
		}
	}

	return ans
}

func helper(s string, l int, r int) string {
	if l < 0 || r > len(s) {
		return ""
	}

	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return s[l+1 : r]
}
