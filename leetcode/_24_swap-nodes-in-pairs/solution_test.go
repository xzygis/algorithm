package _24_swap_nodes_in_pairs

import (
	"testing"

	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{

			name: "test case 1",
			args: args{
				head: SliceToList([]int{1, 2, 3, 4}),
			},
			want: SliceToList([]int{2, 1, 4, 3}),
		},
		{

			name: "test case 2",
			args: args{
				head: nil,
			},
			want: nil,
		},
		{

			name: "test case 3",
			args: args{
				head: SliceToList([]int{1}),
			},
			want: SliceToList([]int{1}),
		},
		{

			name: "test case 4",
			args: args{
				head: SliceToList([]int{1, 2, 3}),
			},
			want: SliceToList([]int{2, 1, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !SliceEquals(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
