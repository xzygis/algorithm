package _24_swap_nodes_in_pairs

import . "github.com/xzygis/algorithm/leetcode/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// A -> B -> C
func swapPairsV1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre := head
	cur := head.Next
	next := head.Next.Next
	cur.Next = pre
	pre.Next = swapPairsV1(next)

	return cur
}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Val: -1, Next: head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}

	return dummyHead.Next
}
