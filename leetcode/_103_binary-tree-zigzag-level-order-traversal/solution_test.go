package _103_binary_tree_zigzag_level_order_traversal

import (
	"reflect"
	"testing"

	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_zigzagLevelOrder(t *testing.T) {
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
				root: ParseLevelTree("[3,9,20,null,null,15,7]"),
			},
			want: ParseMatrix("[[3],[20,9],[15,7]]"),
		},
		{
			name: "test case 2",
			args: args{
				root: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}


