package _234_palindrome_linked_list

import . "github.com/xzygis/algorithm/leetcode/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	tail := tailOfFirstPart(head)
	secondPartHead := reverse(tail.Next)
	for head != nil && secondPartHead != nil {
		if head.Val != secondPartHead.Val {
			return false
		}
		head = head.Next
		secondPartHead = secondPartHead.Next
	}
	return true
}

func tailOfFirstPart(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	return pre
}

func isPalindromeV1(head *ListNode) bool {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	i, j := 0, len(nums)-1
	for i < j {
		if nums[i] != nums[j] {
			return false
		}

		i++
		j--
	}

	return true
}
