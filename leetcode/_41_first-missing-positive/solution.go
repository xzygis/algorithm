package _41_first_missing_positive

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i, num := range nums {
		if num <= 0 {
			nums[i] = n + 1
		}
	}

	for _, num := range nums {
		num = abs(num)
		// 注意此处应该包含等于的情况
		if num <= n {
			nums[num-1] = -abs(nums[num-1])
		}
	}

	for i, num := range nums {
		if num > 0 {
			return i + 1
		}
	}

	return n + 1
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
