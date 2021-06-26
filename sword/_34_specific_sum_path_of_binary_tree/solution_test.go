package _34_specific_sum_path_of_binary_tree

import (
	"reflect"
	"testing"
)

func TestFindPath(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test case 1",
			args: args{
				root:      createTree(),
				targetSum: 22,
			},
			want: [][]int{{10, 5, 7}, {10, 12}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPath(tt.args.root, tt.args.targetSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
		10
	5		12
  4	  7
*/
func createTree() *TreeNode {
	node4 := &TreeNode{4, nil, nil}
	node7 := &TreeNode{7, nil, nil}
	node5 := &TreeNode{5, node4, node7}
	node12 := &TreeNode{12, nil, nil}
	node10 := &TreeNode{10, node5, node12}
	return node10
}
