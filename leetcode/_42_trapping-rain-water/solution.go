package trappingrainwater

func trap(height []int) int {
	leftMax := make([]int, len(height))
	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax := make([]int, len(height))
	rightMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	ans := 0
	for i, h := range height {
		ans += min(leftMax[i], rightMax[i]) - h
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
