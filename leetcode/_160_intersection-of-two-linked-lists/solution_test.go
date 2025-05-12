package _160_intersection_of_two_linked_lists

import (
	"testing"

	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}

	// listA = [4,1,8,4,5], listB = [5,6,1,8,4,5]
	headA := SliceToList([]int{4, 1, 8, 4, 5})
	headB := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: headA.Next}}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{

			name: "test case 1",
			args: args{
				headA: headA,
				headB: headB,
			},
			want: headA.Next,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); got != tt.want {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
