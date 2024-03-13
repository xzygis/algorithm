package longestpalindromicsubstring

func longestPalindrome(s string) string {
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
