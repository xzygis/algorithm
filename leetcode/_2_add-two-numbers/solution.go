package _2_add_two_numbers

import . "github.com/xzygis/algorithm/leetcode/utils"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	cur := result
	carry := 0
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next

		carry = sum / 10
	}

	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}

	return result.Next
}
