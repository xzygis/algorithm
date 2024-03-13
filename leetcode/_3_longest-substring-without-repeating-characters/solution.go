package _3_longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	l := 0
	r := 0
	ans := 0
	for r < len(s) {
		if _, ok := m[s[r]]; ok {
			l = max(l, m[s[r]]+1)
		}

		m[s[r]] = r
		ans = max(ans, r-l+1)
		r++
	}

	return ans
}

func lengthOfLongestSubstringV1(s string) int {
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
