package _198_house_robbe

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	pre2 := nums[0]
	pre1 := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		cur := max(pre2+nums[i], pre1)
		pre2 = pre1
		pre1 = cur
	}

	return pre1
}

// dp[i] = max(dp[i-2] + nums[i], dp[i-1])
func robV1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
