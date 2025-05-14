package _142_linked_list_cycle

import (
	"testing"

	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}

	list1 := SliceToList([]int{3, 2, 0, -4})
	list1.Next.Next.Next.Next = list1.Next

	list2 := SliceToList([]int{1, 2})
	list2.Next.Next = list2

	list3 := SliceToList([]int{1})

	list4 := SliceToList([]int{1, 2})

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{

			name: "test case 1",
			args: args{
				head: list1,
			},
			want: list1.Next,
		},
		{

			name: "test case 2",
			args: args{
				head: list2,
			},
			want: list2,
		},
		{

			name: "test case 3",
			args: args{
				head: list3,
			},
			want: nil,
		},
		{

			name: "test case 4",
			args: args{
				head: list4,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); got != tt.want {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
