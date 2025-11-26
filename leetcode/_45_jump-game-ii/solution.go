package _45_jump_game_ii

func jump(nums []int) int {
	target := len(nums) - 1
	steps := 0
	for target > 0 {
		for i := 0; i < len(nums); i++ {
			if i+nums[i] >= target {
				target = i
				steps++
				break
			}
		}
	}

	return steps
}
