package _25_reverse_nodes_in_k_group

import (
	"testing"

	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{

			name: "test case 1",
			args: args{
				head: SliceToList([]int{1, 2, 3, 4, 5}),
				k:    2,
			},
			want: SliceToList([]int{2, 1, 4, 3, 5}),
		},
		{

			name: "test case 2",
			args: args{
				head: SliceToList([]int{1, 2, 3, 4, 5}),
				k:    3,
			},
			want: SliceToList([]int{3, 2, 1, 4, 5}),
		},
		{

			name: "test case 3",
			args: args{
				head: nil,
				k:    1,
			},
			want: nil,
		},
		{

			name: "test case 4",
			args: args{
				head: SliceToList([]int{1, 2}),
				k:    3,
			},
			want: SliceToList([]int{1, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !SliceEquals(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
