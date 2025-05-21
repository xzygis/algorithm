package _23_merge_k_sorted_lists

import (
	"testing"

	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{

			name: "test case 1",
			args: args{
				lists: []*ListNode{
					SliceToList([]int{1, 4, 5}),
					SliceToList([]int{1, 3, 4}),
					SliceToList([]int{2, 6}),
				},
			},
			want: SliceToList([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{

			name: "test case 2",
			args: args{
				lists: []*ListNode{
					nil,
				},
			},
			want: nil,
		},
		{

			name: "test case 3",
			args: args{
				lists: []*ListNode{
					nil,
					nil,
				},
			},
			want: nil,
		},
		{

			name: "test case 4",
			args: args{
				lists: []*ListNode{
					nil,
					nil,
					nil,
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !SliceEquals(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
