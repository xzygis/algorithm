package _101_symmetric_tree

import . "github.com/xzygis/algorithm/leetcode/utils"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	u, v := root, root
	queue := []*TreeNode{u, v}
	for len(queue) > 0 {
		u, v := queue[0], queue[1]
		queue = queue[2:]
		if u == nil && v == nil {
			continue
		}

		if u == nil || v == nil {
			return false
		}

		if u.Val != v.Val {
			return false
		}

		queue = append(queue, u.Left)
		queue = append(queue, v.Right)

		queue = append(queue, u.Right)
		queue = append(queue, v.Left)
	}

	return true
}

func isSymmetricV1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetric(root.Left, root.Right)
}

func symmetric(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	return symmetric(root1.Left, root2.Right) && symmetric(root1.Right, root2.Left) && root1.Val == root2.Val
}
