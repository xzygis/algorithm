package _113_path_sum_ii

import . "github.com/xzygis/algorithm/leetcode/utils"

func pathSum(root *TreeNode, targetSum int) [][]int {
	var paths [][]int
	var path []int
	helper(root, targetSum, path, &paths)
	return paths
}

func helper(root *TreeNode, targetSum int, path []int, paths *[][]int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		// 注意这里需要深拷贝
		*paths = append(*paths, append([]int{}, path...))
	}

	helper(root.Left, targetSum, path, paths)
	helper(root.Right, targetSum, path, paths)
	path = path[0 : len(path)-1]
}
