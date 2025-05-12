package _160_intersection_of_two_linked_lists

import . "github.com/xzygis/algorithm/leetcode/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pA, pB := headA, headB
	for pA != pB {
		if pA != nil {
			pA = pA.Next
		} else {
			pA = headB
		}

		if pB != nil {
			pB = pB.Next
		} else {
			pB = headA
		}
	}

	return pA
}

func getIntersectionNodeV1(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	for p := headA; p != nil; p = p.Next {
		lenA++
	}

	for p := headB; p != nil; p = p.Next {
		lenB++
	}

	if lenA > lenB {
		for lenA > lenB {
			headA = headA.Next
			lenA--
		}
	}

	if lenB > lenA {
		for lenB > lenA {
			headB = headB.Next
			lenB--
		}
	}

	for headA != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}

	return nil
}
