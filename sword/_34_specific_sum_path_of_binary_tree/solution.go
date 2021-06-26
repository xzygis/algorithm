package _34_specific_sum_path_of_binary_tree

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func FindPath(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}

	var path []int
	var paths [][]int
	curSum := 0
	doFindPaths(root, targetSum, curSum, path, &paths)
	return paths
}

func doFindPaths(root *TreeNode, targetSum int, curSum int, path []int, paths *[][]int) {
	path = append(path, root.val)
	curSum += root.val
	if isLeaf(root) && curSum == targetSum {
		*paths = append(*paths, path)
		fmt.Printf("find path:%v", path)
	}

	if root.left != nil {
		doFindPaths(root.left, targetSum, curSum, path, paths)
	}

	if root.right != nil {
		doFindPaths(root.right, targetSum, curSum, path, paths)
	}
	//返回父节点前，在路径上删除当前节点
	path = path[:len(path)-1]
}

func isLeaf(node *TreeNode) bool {
	return node.left == nil && node.right == nil
}