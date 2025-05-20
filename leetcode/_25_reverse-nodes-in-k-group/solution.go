package _25_reverse_nodes_in_k_group

import . "github.com/xzygis/algorithm/leetcode/utils"

func reverseKGroup(head *ListNode, k int) *ListNode {
	len := 0
	p := head
	for p != nil {
		len++
		p = p.Next
	}

	if len == 0 || len < k {
		return head
	}

	var newHead *ListNode
	cur := head
	for i := 0; i < k; i++ {
		next := cur.Next
		cur.Next = newHead
		newHead = cur
		cur = next
	}

	head.Next = reverseKGroup(cur, k)
	return newHead
}
