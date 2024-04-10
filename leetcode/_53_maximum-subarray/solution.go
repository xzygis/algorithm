package maximumsubarray

func maxSubArray(nums []int) int {
	pre, ans := 0, nums[0]
	for _, num := range nums {
		pre = max(pre+num, num)
		ans = max(ans, pre)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
