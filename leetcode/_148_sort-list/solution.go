package _148_sort_list

import . "github.com/xzygis/algorithm/leetcode/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 注意：如果fast也从head开始，在链表长度为2时，会导致左半部分的长度一直是2
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	mid := slow.Next
	slow.Next = nil

	return mergeList(sortList(head), sortList(mid))
}

func mergeList(head1 *ListNode, head2 *ListNode) *ListNode {
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
