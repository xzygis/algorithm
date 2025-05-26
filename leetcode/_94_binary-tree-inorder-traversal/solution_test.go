package _94_binary_tree_inorder_traversal

import (
	"reflect"
	"testing"

	"github.com/xzygis/algorithm/leetcode/utils"
	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test case 1",
			args: args{
				root: utils.ParseLevelTree("[1,null,2,3]"),
			},
			want: []int{1, 3, 2},
		},
		{
			name: "test case 2",
			args: args{
				root: nil,
			},
			want: nil,
		},
		{
			name: "test case 3",
			args: args{
				root: utils.ParseLevelTree("[1]"),
			},
			want: []int{1},
		},
		{
			name: "test case 4",
			args: args{
				root: utils.ParseLevelTree("[1,2,3,4,5,null,8,null,null,6,7,9]"),
			},
			want: []int{4, 2, 6, 5, 7, 1, 3, 9, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
