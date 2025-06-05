package _105_construct_binary_tree_from_preorder_and_inorder_traversal

import . "github.com/xzygis/algorithm/leetcode/utils"

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
			root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
			break
		}
	}

	return root
}
