package utils

import (
	"strconv"
	"strings"
)

// ParseMatrix 解析二维矩阵，输入格式：[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]
func ParseMatrix(matrixStr string) [][]int {
	matrixStr = strings.TrimPrefix(matrixStr, "[[")
	matrixStr = strings.TrimSuffix(matrixStr, "]]")
	rows := strings.Split(matrixStr, "],[")
	var matrix [][]int
	for _, row := range rows {
		cells := strings.Split(row, ",")
		var nums []int
		for _, cell := range cells {
			num, _ := strconv.Atoi(cell)
			nums = append(nums, num)
		}
		matrix = append(matrix, nums)
	}
	return matrix
}

// ParseMatrixStr 解析二维矩阵，输入格式：[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
func ParseMatrixStr(matrixStr string) [][]int {
	matrixStr = strings.TrimPrefix(matrixStr, "[[")
	matrixStr = strings.TrimSuffix(matrixStr, "]]")
	rows := strings.Split(matrixStr, "],[")
	var matrix [][]int
	for _, row := range rows {
		cells := strings.Split(row, ",")
		var nums []int
		for _, cell := range cells {
			cell = strings.ReplaceAll(cell, "\"", "")
			num, _ := strconv.Atoi(cell)
			nums = append(nums, num)
		}
		matrix = append(matrix, nums)
	}
	return matrix
}