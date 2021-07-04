package _2_add_two_numbers

import (
	. "github.com/xzygis/algorithm/leetcode/utils"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test case 1",
			args: args{
				l1: SliceToList([]int{2, 4, 3}),
				l2: SliceToList([]int{5, 6, 4}),
			},
			want: SliceToList([]int{7, 0, 8}),
		},
		{
			name: "test case 2",
			args: args{
				l1: SliceToList([]int{0}),
				l2: SliceToList([]int{0}),
			},
			want: SliceToList([]int{0}),
		},
		{
			name: "test case 3",
			args: args{
				l1: SliceToList([]int{9, 9, 9, 9, 9, 9, 9}),
				l2: SliceToList([]int{9, 9, 9, 9}),
			},
			want: SliceToList([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !SliceEquals(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
