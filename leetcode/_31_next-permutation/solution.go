package _31_next_permutation

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	// 从后往前找到第一个升序对
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := len(nums) - 1
		// 从后往前找到第一个大于i的数
		for j > i && nums[j] <= nums[i] {
			j--
		}
		// 交换i和j
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 交换整个数组
	swap(nums[i+1:])
}

func swap(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
