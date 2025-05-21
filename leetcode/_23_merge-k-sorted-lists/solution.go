package _23_merge_k_sorted_lists

import . "github.com/xzygis/algorithm/leetcode/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}

	mid := len(lists) / 2
	part1 := mergeKLists(lists[0:mid])
	part2 := mergeKLists(lists[mid:])
	return mergeTwoLists(part1, part2)
}

func mergeTwoLists(head1 *ListNode, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: -1}
	pre := dummyHead
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			pre.Next = head1
			head1 = head1.Next
		} else {
			pre.Next = head2
			head2 = head2.Next
		}

		pre = pre.Next
	}

	if head1 != nil {
		pre.Next = head1
	} else {
		pre.Next = head2
	}

	return dummyHead.Next
}
