package _94_binary_tree_inorder_traversal

import . "github.com/xzygis/algorithm/leetcode/utils"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		ans = append(ans, root.Val)
		stack = stack[0 : len(stack)-1]
		root = root.Right
	}

	return ans
}

func inorderTraversalV1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ans := inorderTraversalV1(root.Left)
	ans = append(ans, root.Val)
	ans = append(ans, inorderTraversalV1(root.Right)...)
	return ans
}
