package _226_invert_binary_tree

import (
	"testing"

	"github.com/xzygis/algorithm/leetcode/utils"
	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_invertTree(t *testing.T) {
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
				root: utils.ParseLevelTree("[4,2,7,1,3,6,9]"),
			},
			want: utils.ParseLevelTree("[4,7,2,9,6,3,1]"),
		},
		{
			name: "test case 2",
			args: args{
				root: utils.ParseLevelTree("[2,1,3]"),
			},
			want: utils.ParseLevelTree("[2,3,1]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !TreeEquals(got, tt.args.root) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
