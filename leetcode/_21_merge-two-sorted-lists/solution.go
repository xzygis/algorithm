package _142_linked_list_cycle

import . "github.com/xzygis/algorithm/leetcode/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoListsV1(list1 *ListNode, list2 *ListNode) *ListNode {
	var newHead = &ListNode{Val: -1}
	var pre = newHead
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			pre.Next = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
		}

		pre = pre.Next
	}

	if list1 == nil {
		pre.Next = list2
	} else {
		pre.Next = list1
	}

	return newHead.Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}
