package _19_remove_nth_node_from_end_of_list

import (
	. "github.com/xzygis/algorithm/leetcode/utils"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEndV1(head *ListNode, n int) *ListNode {
	len := 0
	p := head
	for p != nil {
		len++
		p = p.Next
	}

	var dummy = &ListNode{Val: -1, Next: head}
	pre := dummy
	for i := 0; i < len-n; i++ {
		pre = pre.Next
	}

	pre.Next = pre.Next.Next
	return dummy.Next
}

/*
0    ->    1    ->    2    ->    3    ->    4    ->    5    ->    nil
n=2
s          f
s                                f
                                 s                                 f
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	first := head //注意此处为head
	second := dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next
	return dummy.Next
}
