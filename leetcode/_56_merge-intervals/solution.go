package _56_merge_intervals

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var ans [][]int
	ans = append(ans, intervals[0])

	for i := 1; i < len(intervals); i++ {
		left := intervals[i][0]
		right := intervals[i][1]
		lastRight := ans[len(ans)-1][1]
		if left <= lastRight {
			if right > lastRight {
				ans[len(ans)-1][1] = right
			}
		} else {
			ans = append(ans, intervals[i])
		}
	}

	return ans
}
