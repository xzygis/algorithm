package _55_jump_game

func canJump(nums []int) bool {
	rightMost := 0
	for i := 0; i < len(nums); i++ {
		if i <= rightMost {
			// 注意：必须包含=的情况，特殊case是输入为：[0]
			if i+nums[i] >= rightMost {
				rightMost = i + nums[i]
				if rightMost >= len(nums)-1 {
					return true
				}
			}
		}
	}

	return false
}

func canJumpV1(nums []int) bool {
	return helper(nums, 0)
}

func helper(nums []int, start int) bool {
	if start == len(nums)-1 {
		return true
	}

	if nums[start] == 0 {
		return false
	}

	for i := 1; i <= nums[start]; i++ {
		if helper(nums, start+i) {
			return true
		}
	}

	return false
}
