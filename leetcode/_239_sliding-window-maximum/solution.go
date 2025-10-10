package _239_sliding_window_maximum

func maxSlidingWindow(nums []int, k int) []int {
	var queue []int
	var push = func(i int) {
		for len(queue) > 0 && nums[i] > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	var ans []int
	ans = append(ans, nums[queue[0]])
	for i := k; i < len(nums); i++ {
		push(i)
		for len(queue) > 0 && queue[0] <= i-k {
			queue = queue[1:]
		}
		ans = append(ans, nums[queue[0]])
	}

	return ans
}
