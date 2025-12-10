package _152_maximum_product_subarray

func maxProduct(nums []int) int {
	preMin := nums[0]
	preMax := nums[0]
	ans := preMin
	for i := 1; i < len(nums); i++ {
		// 如果nums[i]为负数，最大值来自于minDp[i-1]*nums[i]
		if nums[i] < 0 {
			preMin, preMax = preMax, preMin
		}
		preMin = min(preMin*nums[i], nums[i])
		preMax = max(preMax*nums[i], nums[i])
		ans = max(ans, preMax)
	}

	return ans
}

func maxProductV1(nums []int) int {
	minDp := make([]int, len(nums))
	maxDp := make([]int, len(nums))
	minDp[0] = nums[0]
	maxDp[0] = nums[0]
	ans := maxDp[0]
	for i := 1; i < len(nums); i++ {
		// 如果nums[i]为负数，最大值来自于minDp[i-1]*nums[i]
		minDp[i] = min(nums[i], min(minDp[i-1]*nums[i], maxDp[i-1]*nums[i]))
		maxDp[i] = max(nums[i], max(minDp[i-1]*nums[i], maxDp[i-1]*nums[i]))
		ans = max(ans, maxDp[i])
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
