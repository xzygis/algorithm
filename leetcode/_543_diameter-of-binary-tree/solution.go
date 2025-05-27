package _543_diameter_of_binary_tree

import . "github.com/xzygis/algorithm/leetcode/utils"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	depth(root, &ans)
	return ans - 1
}

func depth(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}

	l := depth(root.Left, ans)
	r := depth(root.Right, ans)
	*ans = max(*ans, l+r+1)
	return max(l, r) + 1
}

func diameterOfBinaryTreeV1(root *TreeNode) int {
	path, _ := helper(root)
	return path - 1
}

func helper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	lpath, ldepth := helper(root.Left)
	rpath, rdepth := helper(root.Right)
	return max(max(lpath, rpath), ldepth+rdepth+1), max(ldepth, rdepth) + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
