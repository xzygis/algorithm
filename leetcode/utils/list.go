package utils

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func SliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	result := &ListNode{Val: nums[0]}
	pre := result
	for i := 1; i < len(nums); i++ {
		pre.Next = &ListNode{Val: nums[i]}
		pre = pre.Next
	}
	return result
}

func SliceEquals(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil || l2 != nil {
		if l1 == nil || l2 == nil {
			return false
		}

		if l1.Val != l2.Val {
			return false
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	return true
}

func PrintList(head *ListNode) {
	sb := strings.Builder{}
	for head != nil {
		sb.WriteString(fmt.Sprintf("%v ", head.Val))
		head = head.Next
	}
	println(sb.String())
}
