package _46_permutations

func permute(nums []int) [][]int {
	var ans [][]int
	backtrack(nums, 0, &ans)
	return ans
}

func backtrack(nums []int, start int, ans *[][]int) {
	if start == len(nums)-1 {
		*ans = append(*ans, append([]int{}, nums...))
		return
	}

	// 从start开始，依次交换元素生成排列
	for i := start; i < len(nums); i++ {
		nums[i], nums[start] = nums[start], nums[i]
		backtrack(nums, start+1, ans)
		// 回溯，恢复原始顺序
		nums[i], nums[start] = nums[start], nums[i]
	}
}
