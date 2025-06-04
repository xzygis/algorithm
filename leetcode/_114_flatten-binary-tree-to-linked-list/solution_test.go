package _114_flatten_binary_tree_to_linked_list

import (
	"testing"

	"github.com/xzygis/algorithm/leetcode/utils"
	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_flatten(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "test case 1",
			args: args{
				root: utils.ParseLevelTree("[1,2,5,3,4,null,6]"),
			},
			want: utils.ParseLevelTree("[1,null,2,null,3,null,4,null,5,null,6]"),
		},
		{
			name: "test case 2",
			args: args{
				root: utils.ParseLevelTree("[]"),
			},
			want: utils.ParseLevelTree("[]"),
		},
		{
			name: "test case 3",
			args: args{
				root: utils.ParseLevelTree("[0]"),
			},
			want: utils.ParseLevelTree("[0]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if flatten(tt.args.root); !TreeEquals(tt.args.root, tt.want) {
				utils.PrintTree(tt.args.root)
				utils.PrintTree(tt.want)
				t.Errorf("flatten() = %v, want %v", tt.args.root, tt.want)
			}
		})
	}
}
