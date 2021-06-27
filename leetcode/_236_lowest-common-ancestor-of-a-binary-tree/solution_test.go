package _236_lowest_common_ancestor_of_a_binary_tree

import (
	"reflect"
	"testing"

	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}

	root := ParseLevelTree("[3,5,1,6,2,0,8,null,null,7,4]")

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "test case 1",
			args: args{
				root: root,
				p:    root.Left,
				q:    root.Right,
			},
			want: root,
		},
		{
			name: "test case 2",
			args: args{
				root: root,
				p:    root.Left,
				q:    root.Left.Right.Right,
			},
			want: root.Left,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
