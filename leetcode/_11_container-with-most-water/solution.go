package containerwithmostwater

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left <= right {
		area := min(height[left], height[right]) * (right - left)
		maxArea = max(maxArea, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
