package _46_permutations

import (
	"reflect"
	"sort"
	"testing"
)

func TestPermute(t *testing.T) {
	testCases := []struct {
		input  []int
		output [][]int
	}{
		{
			input: []int{1, 2, 3},
			output: [][]int{
				{1, 2, 3}, {1, 3, 2}, {2, 1, 3},
				{2, 3, 1}, {3, 1, 2}, {3, 2, 1},
			},
		},
		{
			input:  []int{0, 1},
			output: [][]int{{0, 1}, {1, 0}},
		},
		{
			input:  []int{1},
			output: [][]int{{1}},
		},
	}

	for _, tc := range testCases {
		result := permute(tc.input)
		if !equalPermutations(result, tc.output) {
			t.Errorf("permute(%v) = %v, want %v",
				tc.input, result, tc.output)
		}
	}
}

// 辅助函数：比较两个排列集合是否相等（忽略顺序）
func equalPermutations(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// 对每个排列进行排序
	sortSlice := func(slice [][]int) {
		sort.Slice(slice, func(i, j int) bool {
			for k := 0; k < len(slice[i]); k++ {
				if slice[i][k] != slice[j][k] {
					return slice[i][k] < slice[j][k]
				}
			}
			return false
		})
	}

	// 复制以避免修改原数组
	aCopy := make([][]int, len(a))
	bCopy := make([][]int, len(b))
	for i := range a {
		aCopy[i] = make([]int, len(a[i]))
		copy(aCopy[i], a[i])
		bCopy[i] = make([]int, len(b[i]))
		copy(bCopy[i], b[i])
	}

	sortSlice(aCopy)
	sortSlice(bCopy)

	return reflect.DeepEqual(aCopy, bCopy)
}
