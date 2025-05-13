package _234_palindrome_linked_list

import (
	"testing"

	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{

			name: "test case 1",
			args: args{
				head: SliceToList([]int{1, 2, 2, 1}),
			},
			want: true,
		},
		{

			name: "test case 1",
			args: args{
				head: SliceToList([]int{1, 2}),
			},
			want: false,
		},
		{

			name: "test case 1",
			args: args{
				head: SliceToList([]int{1}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
