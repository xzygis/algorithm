package _112_path_sum

import (
	"reflect"
	"testing"

	"github.com/xzygis/algorithm/leetcode/utils"
	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root   *TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case 1",
			args: args{
				root:   utils.ParseLevelTree("[5,4,8,11,null,13,4,7,2,null,null,null,1]"),
				target: 22,
			},
			want: true,
		},
		{
			name: "test case 2",
			args: args{
				root:   utils.ParseLevelTree("[1,2,3]"),
				target: 5,
			},
			want: false,
		},
		{
			name: "test case 3",
			args: args{
				root:   utils.ParseLevelTree("[]"),
				target: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
