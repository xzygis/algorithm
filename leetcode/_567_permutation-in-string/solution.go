package permutationinstring

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	count := [26]int{}
	for _, ch := range s1 {
		count[ch-'a']--
	}

	left := 0
	for right, ch := range s2 {
		count[ch-'a']++
		for count[ch-'a'] > 0 {
			count[s2[left]-'a']--
			left++
		}

		if right-left+1 == len(s1) {
			return true
		}
	}

	return false
}
