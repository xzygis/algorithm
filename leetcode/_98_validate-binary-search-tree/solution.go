package _98_validate_binary_search_tree

import (
	"math"

	. "github.com/xzygis/algorithm/leetcode/utils"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	var q []*TreeNode
	var last = math.MinInt
	for root != nil || len(q) > 0 {
		for root != nil {
			q = append(q, root)
			root = root.Left
		}

		root = q[len(q)-1]
		q = q[0 : len(q)-1]
		if root.Val <= last {
			return false
		}
		last = root.Val
		root = root.Right
	}

	return true
}

func isValidBSTV1(root *TreeNode) bool {
	return helper(root, math.MinInt, math.MaxInt)
}

func helper(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)
}
