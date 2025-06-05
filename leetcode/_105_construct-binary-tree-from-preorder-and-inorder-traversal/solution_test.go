package _105_construct_binary_tree_from_preorder_and_inorder_traversal

import (
	"testing"

	"github.com/xzygis/algorithm/leetcode/utils"
	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "test case 1",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: utils.ParseLevelTree("[3,9,20,null,null,15,7]"),
		},
		{
			name: "test case 2",
			args: args{
				preorder: []int{-1},
				inorder:  []int{-1},
			},
			want: utils.ParseLevelTree("[-1]"),
		},
		{
			name: "test case 3",
			args: args{
				preorder: []int{1, 2},
				inorder:  []int{2, 1},
			},
			want: utils.ParseLevelTree("[1,2]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !utils.TreeEquals(got, tt.want) {
				utils.PrintTree(got)
				utils.PrintTree(tt.want)
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
