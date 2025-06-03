package _102_binary_tree_level_order_traversal

import (
	"reflect"
	"testing"

	"github.com/xzygis/algorithm/leetcode/utils"
	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test case 1",
			args: args{
				root: utils.ParseLevelTree("[3,9,20,null,null,15,7]"),
			},
			want: utils.ParseMatrix("[[3],[9,20],[15,7]]"),
		},
		{
			name: "test case 2",
			args: args{
				root: utils.ParseLevelTree("[1]"),
			},
			want: utils.ParseMatrix("[[1]]"),
		},
		{
			name: "test case 3",
			args: args{
				root: utils.ParseLevelTree("[]"),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
