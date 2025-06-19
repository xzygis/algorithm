package _47_permutations_ii

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	var path []int
	var vis = make([]bool, len(nums))
	backtrack(nums, 0, vis, path, &ans)
	return ans
}

func backtrack(nums []int, start int, vis []bool, path []int, ans *[][]int) {
	if len(path) == len(nums) {
		*ans = append(*ans, append([]int{}, path...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if vis[i] || (i > 0 && nums[i] == nums[i-1] && !vis[i-1]) {
			continue
		}

		vis[i] = true
		path = append(path, nums[i])
		backtrack(nums, start+1, vis, path, ans)
		vis[i] = false
		path = path[0 : len(path)-1]
	}
}
