package _236_lowest_common_ancestor_of_a_binary_tree

import (
	. "github.com/xzygis/algorithm/leetcode/utils"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	if p == nil {
		return q
	}

	if q == nil {
		return p
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else {
		return right
	}
}

