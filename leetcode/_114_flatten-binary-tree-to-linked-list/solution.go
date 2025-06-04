package _114_flatten_binary_tree_to_linked_list

import . "github.com/xzygis/algorithm/leetcode/utils"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	stack := []*TreeNode{root}
	var pre *TreeNode
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		if pre != nil {
			pre.Left = nil
			pre.Right = cur
		}

		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}

		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}

		pre = cur
	}
}

func flattenV2(root *TreeNode) {
	nodes := helperV2(root)
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Right = nodes[i+1]
		nodes[i].Left = nil
	}
}

func helperV2(root *TreeNode) []*TreeNode {
	var stack []*TreeNode
	var nodes []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			nodes = append(nodes, root)
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		root = root.Right
	}
	return nodes
}

func flattenV1(root *TreeNode) {
	nodes := helperV1(root)
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Right = nodes[i+1]
		nodes[i].Left = nil
	}
}

func helperV1(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	nodes := []*TreeNode{root}
	nodes = append(nodes, helperV1(root.Left)...)
	nodes = append(nodes, helperV1(root.Right)...)
	return nodes
}
