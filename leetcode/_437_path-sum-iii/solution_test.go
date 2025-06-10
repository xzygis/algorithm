package _437_path_sum_iii

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
		want int
	}{
		{
			name: "test case 1",
			args: args{
				root:   utils.ParseLevelTree("[10,5,-3,3,2,null,11,3,-2,null,1]"),
				target: 8,
			},
			want: 3,
		},
		{
			name: "test case 2",
			args: args{
				root:   utils.ParseLevelTree("[5,4,8,11,null,13,4,7,2,null,null,5,1]"),
				target: 22,
			},
			want: 3,
		},
		{
			name: "test case 3",
			args: args{
				root:   utils.ParseLevelTree("[]"),
				target: 0,
			},
			want: 0,
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
