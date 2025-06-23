package _78_subsets

import (
	"sort"
	"strconv"
	"testing"
)

func TestSubsets(t *testing.T) {
	type TestCase struct {
		input    []int
		expected [][]int
	}
	testCases := []TestCase{
		{
			input:    []int{1, 2, 3},
			expected: [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			input:    []int{0},
			expected: [][]int{{}, {0}},
		},
	}
	for _, tc := range testCases {
		result := subsets(tc.input)
		// 分别对结果和期望输出进行排序处理
		sortedResult := sortSubsets(result)
		sortedExpected := sortSubsets(tc.expected)
		// 序列化为字符串比较
		if serialize(sortedResult) != serialize(sortedExpected) {
			t.Errorf("输入: %v, 期望输出: %v, 实际输出: %v", tc.input, tc.expected, result)
		}
	}
}

// 对切片的切片进行排序，先按长度，再按元素值
func sortSubsets(subsets [][]int) [][]int {
	sort.Slice(subsets, func(i, j int) bool {
		if len(subsets[i]) != len(subsets[j]) {
			return len(subsets[i]) < len(subsets[j])
		}
		for k := range subsets[i] {
			if subsets[i][k] != subsets[j][k] {
				return subsets[i][k] < subsets[j][k]
			}
		}
		return false
	})
	return subsets
}

// 序列化为字符串，方便比较
func serialize(subsets [][]int) string {
	var s string
	for _, sub := range subsets {
		for _, num := range sub {
			s += strconv.Itoa(num)
		}
		s += "|"
	}
	return s
}
