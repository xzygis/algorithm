package _437_path_sum_iii

import . "github.com/xzygis/algorithm/leetcode/utils"

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	ans := rootSum(root, targetSum)
	ans += pathSum(root.Left, targetSum)
	ans += pathSum(root.Right, targetSum)
	return ans
}

func rootSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	ans := 0
	if root.Val == targetSum {
		ans++
	}

	ans += rootSum(root.Left, targetSum-root.Val)
	ans += rootSum(root.Right, targetSum-root.Val)
	return ans
}
