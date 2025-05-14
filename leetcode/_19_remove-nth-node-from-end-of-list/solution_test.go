package _19_remove_nth_node_from_end_of_list

import (
	"testing"

	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
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
				n:    2,
			},
			want: SliceToList([]int{1, 2, 3, 5}),
		},
		{

			name: "test case 2",
			args: args{
				head: SliceToList([]int{1}),
				n:    1,
			},
			want: nil,
		},
		{

			name: "test case 3",
			args: args{
				head: SliceToList([]int{1, 2}),
				n:    1,
			},
			want: SliceToList([]int{1}),
		},
		{

			name: "test case 4",
			args: args{
				head: SliceToList([]int{1, 2}),
				n:    2,
			},
			want: SliceToList([]int{2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !SliceEquals(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
