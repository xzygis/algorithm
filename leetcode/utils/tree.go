package utils

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// [3,9,20,null,null,15,7]
func ParseLevelTree(treeStr string) *TreeNode {
	treeStr = strings.TrimPrefix(treeStr, "[")
	treeStr = strings.TrimSuffix(treeStr, "]")
	if treeStr == "" {
		return nil
	}

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
