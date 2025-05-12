package _206_reverse_linked_list

import . "github.com/xzygis/algorithm/leetcode/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	// 为避免空指针问题，此处应该包含head.Next == nil的情况
	if head == nil || head.Next == nil {
		return head
	}

	preHead := head
	cur := head.Next
	for cur != nil {
		next := cur.Next //暂存下一个节点
		cur.Next = preHead
		preHead = cur
		cur = next
	}
	head.Next = nil
	return preHead
}

func reverseListV1(head *ListNode) *ListNode {
	// 为避免空指针问题，此处应该包含head.Next == nil的情况
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
