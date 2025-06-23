package _78_subsets

func subsets(nums []int) [][]int {
	var ans [][]int
	for mask := 0; mask < 1<<len(nums); mask++ {
		var set []int
		for i := 0; i < len(nums); i++ {
			if mask>>i&1 == 1 {
				set = append(set, nums[i])
			}
		}
		ans = append(ans, set)
	}
	return ans
}

func subsetsV1(nums []int) [][]int {
	var ans [][]int                // 存储最终的所有子集
	var path []int                 // 存储当前递归路径上的元素
	backtrace(nums, 0, path, &ans) // 从索引 0 开始回溯
	return ans
}

func backtrace(nums []int, start int, path []int, ans *[][]int) {
	// 将当前路径的深拷贝加入结果集（避免后续修改影响已加入的路径）
	*ans = append(*ans, append([]int{}, path...))

	// 从 start 开始遍历剩余元素，避免重复组合
	for i := start; i < len(nums); i++ {
		// 选择当前元素
		path = append(path, nums[i])
		// 递归处理下一个元素（注意这里传入的是 start+1，导致错误）
		backtrace(nums, i+1, path, ans)
		// 回溯：撤销选择，移除最后一个元素，尝试其他可能性
		path = path[0 : len(path)-1]
	}
}
