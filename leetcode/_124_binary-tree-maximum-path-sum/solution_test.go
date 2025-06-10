package _124_binary_tree_maximum_path_sum

import (
	"math"
	"reflect"
	"testing"

	"github.com/xzygis/algorithm/leetcode/utils"
	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_maxPathSum(t *testing.T) {
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
				root: utils.ParseLevelTree("[1,2,3]"),
			},
			want: 6,
		},
		{
			name: "test case 2",
			args: args{
				root: utils.ParseLevelTree("[-10,9,20,null,null,15,7]"),
			},
			want: 42,
		},
		{
			name: "test case 3",
			args: args{
				root: utils.ParseLevelTree("[]"),
			},
			want: math.MinInt,
		},
		{
			name: "test case 4",
			args: args{
				root: utils.ParseLevelTree("[-3]"),
			},
			want: -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
