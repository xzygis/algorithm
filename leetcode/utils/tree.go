package utils

import (
	"fmt"
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

func PrintNodes(nodes []*TreeNode) {
	sb := strings.Builder{}
	for _, node := range nodes {
		sb.WriteString(fmt.Sprintf("%v ", node.Val))
	}
	println(sb.String())
}

func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}

	sb := strings.Builder{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			sb.WriteString(fmt.Sprintf("%v ", node.Val))
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			size--
		}
		sb.WriteString("\n")
	}

	println(sb.String())
}

func TreeEquals(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil {
		return false
	}

	if root2 == nil {
		return false
	}

	return TreeEquals(root1.Left, root2.Left) && TreeEquals(root1.Right, root2.Right) && root1.Val == root2.Val
}
