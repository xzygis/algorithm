package _113_path_sum_ii

import (
	"reflect"
	"testing"

	"github.com/xzygis/algorithm/leetcode/utils"
	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root   *TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test case 1",
			args: args{
				root:   utils.ParseLevelTree("[5,4,8,11,null,13,4,7,2,null,null,5,1]"),
				target: 22,
			},
			want: utils.ParseMatrix("[[5,4,11,2],[5,8,4,5]]"),
		},
		{
			name: "test case 2",
			args: args{
				root:   utils.ParseLevelTree("[1,2,3]"),
				target: 5,
			},
			want: nil,
		},
		{
			name: "test case 3",
			args: args{
				root:   utils.ParseLevelTree("[1,2]"),
				target: 0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
