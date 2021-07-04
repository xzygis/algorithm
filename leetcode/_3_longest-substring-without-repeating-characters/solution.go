package _3_longest_substring_without_repeating_characters

// abcabcbb
// dvdf
// abba
func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	l := 0
	r := -1
	ans := 0
	for l < len(s) {
		if l != 0 {
			delete(m, s[l-1])
		}

		for r+1 < len(s) && m[s[r+1]] == 0 {
			m[s[r+1]]++
			r++
		}

		ans = max(ans, r-l+1)
		l++
	}

	return ans
}

func max(m, n int) int {
	if m > n {
		return m
	}

	return n
}
