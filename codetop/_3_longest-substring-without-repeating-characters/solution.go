package _3_longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	m := make(map[uint8]int)
	left, right := 0, 0
	ans := 0
	for right < len(s) {
		ch := s[right]
		m[ch]++
		for m[ch] > 1 {
			m[s[left]]--
			left++
		}

		ans = max(ans, right-left+1)
		right++
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
