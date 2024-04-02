package _438_find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	var res []int
	var pCount [26]int
	var sCount [26]int
	for i, ch := range p {
		pCount[ch-'a']++
		sCount[s[i]-'a']++
	}

	if pCount == sCount {
		res = append(res, 0)
	}

	for i, ch := range s[:len(s)-len(p)] {
		sCount[ch-'a']--
		sCount[s[i+len(p)]-'a']++

		if sCount == pCount {
			res = append(res, i+1)
		}
	}

	return res
}
