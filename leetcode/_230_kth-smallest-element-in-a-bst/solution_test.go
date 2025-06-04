package _230_kth_smallest_element_in_a_bst

import (
	"reflect"
	"testing"

	"github.com/xzygis/algorithm/leetcode/utils"
	. "github.com/xzygis/algorithm/leetcode/utils"
)

func Test_kthSmallest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				root: utils.ParseLevelTree("[3,1,4,null,2]"),
				k:    1,
			},
			want: 1,
		},
		{
			name: "test case 2",
			args: args{
				root: utils.ParseLevelTree("[5,3,6,2,4,null,null,1]"),
				k:    3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.root, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
