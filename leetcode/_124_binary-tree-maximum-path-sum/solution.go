package _124_binary_tree_maximum_path_sum

import (
	"math"

	. "github.com/xzygis/algorithm/leetcode/utils"
)

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt
	var maxGain func(root *TreeNode) int
	maxGain = func(root *TreeNode) int {
		if root == nil {
			// impossible case
			return 0
		}

		leftGain := max(maxGain(root.Left), 0)
		rightGain := max(maxGain(root.Right), 0)

		maxSum = max(maxSum, leftGain+root.Val+rightGain)
		return root.Val + max(leftGain, rightGain)
	}

	maxGain(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
