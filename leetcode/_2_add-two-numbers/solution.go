package _2_add_two_numbers

import . "github.com/xzygis/algorithm/leetcode/utils"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var newHead = &ListNode{Val: -1}
	var pre = newHead
	carry := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		pre.Next = &ListNode{Val: carry % 10}
		pre = pre.Next

		carry /= 10
	}

	if carry > 0 {
		pre.Next = &ListNode{Val: carry}
	}

	return newHead.Next
}
