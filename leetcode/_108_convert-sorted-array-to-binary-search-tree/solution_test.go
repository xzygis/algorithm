package _108_convert_sorted_array_to_binary_search_tree

import (
	"testing"

	"github.com/xzygis/algorithm/leetcode/utils"
	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "test case 1",
			args: args{
				nums: []int{-10, -3, 0, 5, 9},
			},
			want: utils.ParseLevelTree("[0,-3,9,-10,null,5]"),
		},
		{
			name: "test case 2",
			args: args{
				nums: []int{1, 3},
			},
			want: utils.ParseLevelTree("[3, 1]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !utils.TreeEquals(got, tt.want) {
				utils.PrintTree(got)
				utils.PrintTree(tt.want)
				utils.PrintTree(utils.ParseLevelTree("[3, 1]"))
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
