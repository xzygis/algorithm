package _763_partition_labels

func partitionLabels(s string) []int {
	lastIndex := make(map[byte]int)
	for i, ch := range s {
		lastIndex[byte(ch)] = i
	}

	start, end := 0, 0
	var ans []int
	for i := 0; i < len(s); i++ {
		if lastIndex[s[i]] > end {
			end = lastIndex[s[i]]
		}

		if i == end {
			ans = append(ans, end-start+1)
			start = i + 1
		}
	}

	return ans
}
