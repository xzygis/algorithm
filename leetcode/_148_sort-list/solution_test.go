package _148_sort_list

import (
	"testing"

	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_sortList(t *testing.T) {
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
				head: SliceToList([]int{4, 2, 1, 3}),
			},
			want: SliceToList([]int{1, 2, 3, 4}),
		},
		{

			name: "test case 2",
			args: args{
				head: SliceToList([]int{-1, 5, 3, 4, 0}),
			},
			want: SliceToList([]int{-1, 0, 3, 4, 5}),
		},
		{

			name: "test case 3",
			args: args{
				head: nil,
			},
			want: nil,
		},
		{

			name: "test case 4",
			args: args{
				head: SliceToList([]int{-1, 5, 3, 4, 0}),
			},
			want: SliceToList([]int{-1, 0, 3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !SliceEquals(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
