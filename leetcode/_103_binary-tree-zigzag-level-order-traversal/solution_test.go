package _103_binary_tree_zigzag_level_order_traversal

import (
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/xzygis/algorithm/leetcode/utils"
)

func Test_zigzagLevelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test case 1",
			args: args{
				root: parseLevelTree("[3,9,20,null,null,15,7]"),
			},
			want: utils.ParseMatrix("[[3],[20,9],[15,7]]"),
		},
		{
			name: "test case 2",
			args: args{
				root: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

//[3,9,20,null,null,15,7]
func parseLevelTree(treeStr string) *TreeNode {
	treeStr = strings.TrimPrefix(treeStr, "[")
	treeStr = strings.TrimSuffix(treeStr, "]")
	vals := strings.Split(treeStr, ",")
	nodes := createNodes(vals)
	return createTree(nodes)
}

func createTree(nodes []*TreeNode) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}

	root := nodes[0]
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if i < len(nodes) {
			node.Left = nodes[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}

		if i+1 < len(nodes) {
			node.Right = nodes[i+1]
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		i += 2
	}
	return root
}

func createNodes(vals []string) []*TreeNode {
	nodes := make([]*TreeNode, 0, len(vals))
	for _, val := range vals {
		if val == "null" {
			nodes = append(nodes, nil)
		} else {
			num, _ := strconv.Atoi(val)
			nodes = append(nodes, &TreeNode{Val: num})
		}
	}
	return nodes
}
