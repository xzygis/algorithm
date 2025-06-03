package _108_convert_sorted_array_to_binary_search_tree

import . "github.com/xzygis/algorithm/leetcode/utils"

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[0:mid])
	if mid < len(nums) {
		root.Right = sortedArrayToBST(nums[mid+1:])
	}

	return root
}
