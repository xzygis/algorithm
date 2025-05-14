package _142_linked_list_cycle

import . "github.com/xzygis/algorithm/leetcode/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			// 第一次相遇
			// 让p指向头结点，与慢指针一起往后走，直到相遇
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}

			return p
		}
	}

	return nil
}
