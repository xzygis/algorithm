package binarytreerightsideview

import (
	"reflect"
	"testing"

	"github.com/xzygis/algorithm/leetcode/utils"
	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_rightSideView(t *testing.T) {
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
				root: utils.ParseLevelTree("[1,2,3,null,5,null,4]"),
			},
			want: []int{1, 3, 4},
		},
		{
			name: "test case 2",
			args: args{
				root: utils.ParseLevelTree("[1,null,3]"),
			},
			want: []int{1, 3},
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
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
