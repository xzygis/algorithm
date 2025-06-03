package _102_binary_tree_level_order_traversal

import . "github.com/xzygis/algorithm/leetcode/utils"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ans [][]int
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		var row []int
		for size > 0 {
			node := q[0]
			q = q[1:]
			row = append(row, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			size--
		}
		ans = append(ans, row)
	}

	return ans
}
