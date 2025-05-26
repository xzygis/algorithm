package _104_maximum_depth_of_binary_tree

import (
	"testing"

	"github.com/xzygis/algorithm/leetcode/utils"
	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				root: utils.ParseLevelTree("[3,9,20,null,null,15,7]"),
			},
			want: 3,
		},
		{
			name: "test case 2",
			args: args{
				root: utils.ParseLevelTree("[1,null,2]"),
			},
			want: 2,
		},
		{
			name: "test case 3",
			args: args{
				root: utils.ParseLevelTree("[0,2,4,1,null,3,-1,5,1,null,6,null,8]"),
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
