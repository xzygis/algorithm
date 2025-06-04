package _230_kth_smallest_element_in_a_bst

import . "github.com/xzygis/algorithm/leetcode/utils"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	i := 0
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		i++
		if i == k {
			return root.Val
		}

		root = root.Right
	}

	return -1
}
