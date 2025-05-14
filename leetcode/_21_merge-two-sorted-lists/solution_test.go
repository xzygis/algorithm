package _142_linked_list_cycle

import (
	"testing"

	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{

			name: "test case 1",
			args: args{
				list1: SliceToList([]int{1, 2, 4}),
				list2: SliceToList([]int{1, 3, 4}),
			},
			want: SliceToList([]int{1, 1, 2, 3, 4, 4}),
		},
		{

			name: "test case 2",
			args: args{
				list1: nil,
				list2: nil,
			},
			want: nil,
		},
		{

			name: "test case 3",
			args: args{
				list1: SliceToList([]int{0}),
				list2: nil,
			},
			want: SliceToList([]int{0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !SliceEquals(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
