package _103_binary_tree_zigzag_level_order_traversal

import (
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	reverse := false
	var result [][]int
	for len(queue) > 0 {
		size := len(queue)
		row := make([]int, size, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if reverse {
				row[size-i-1] = node.Val
			} else {
				row[i] = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, row)
		reverse = !reverse
	}
	return result
}
