package longestconsecutivesequence

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}

	maxCount := 0
	for _, num := range nums {
		if !numSet[num-1] {
			curCount := 0
			for numSet[num] {
				curCount++
				num++
			}

			if curCount > maxCount {
				maxCount = curCount
			}
		}
	}

	return maxCount
}
