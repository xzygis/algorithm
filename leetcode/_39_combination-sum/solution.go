package _39_combination_sum

/*
输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
*/
func combinationSum(candidates []int, target int) [][]int {
	var path []int
	var ans [][]int
	backtrack(candidates, target, 0, path, &ans)
	return ans
}

func backtrack(candidates []int, target int, start int, path []int, ans *[][]int) {
	if target == 0 {
		*ans = append(*ans, append([]int{}, path...))
		return
	}

	if target < 0 {
		return
	}

	for i := start; i < len(candidates); i++ {
		path = append(path, candidates[i])
		backtrack(candidates, target-candidates[i], i, path, ans)
		path = path[0 : len(path)-1]
	}
}
