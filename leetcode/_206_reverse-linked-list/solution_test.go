package _206_reverse_linked_list

import (
	"testing"

	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_reverseList(t *testing.T) {
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
				head: SliceToList([]int{1, 2, 3, 4, 5}),
			},
			want: SliceToList([]int{5, 4, 3, 2, 1}),
		},
		{

			name: "test case 2",
			args: args{
				head: SliceToList([]int{1, 2}),
			},
			want: SliceToList([]int{2, 1}),
		},
		{

			name: "test case 3",
			args: args{
				head: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !SliceEquals(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
