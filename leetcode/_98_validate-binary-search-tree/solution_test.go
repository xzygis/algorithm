package _98_validate_binary_search_tree

import (
	"reflect"
	"testing"

	"github.com/xzygis/algorithm/leetcode/utils"
	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case 1",
			args: args{
				root: utils.ParseLevelTree("[2,1,3]"),
			},
			want: true,
		},
		{
			name: "test case 2",
			args: args{
				root: utils.ParseLevelTree("[5,1,4,null,null,3,6]"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
