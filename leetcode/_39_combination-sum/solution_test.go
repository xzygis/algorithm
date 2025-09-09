package _39_combination_sum

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			name:       "Example1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expected:   [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "Example2",
			candidates: []int{2, 3, 5},
			target:     8,
			expected:   [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "Example3",
			candidates: []int{2},
			target:     1,
			expected:   [][]int{},
		},
		{
			name:       "SingleElement",
			candidates: []int{3},
			target:     6,
			expected:   [][]int{{3, 3}},
		},
		{
			name:       "EmptyInput",
			candidates: []int{},
			target:     5,
			expected:   [][]int{},
		},
		{
			name:       "NoSolution",
			candidates: []int{3, 6, 9},
			target:     5,
			expected:   [][]int{},
		},
		{
			name:       "LargeNumbers",
			candidates: []int{10, 20, 30},
			target:     50,
			expected: [][]int{
				{10, 10, 10, 10, 10},
				{10, 10, 30},
				{10, 20, 20},
				{10, 10, 10, 20},
				{20, 30},
			},
		},
		{
			name:       "DuplicatePaths",
			candidates: []int{2, 4},
			target:     6,
			expected:   [][]int{{2, 2, 2}, {2, 4}},
		},
		{
			name:       "TargetZero",
			candidates: []int{1, 2, 3},
			target:     0,
			expected:   [][]int{{}}, // 注意: 空组合是唯一解
		},
		{
			name:       "TestWithOne",
			candidates: []int{1},
			target:     2,
			expected:   [][]int{{1, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := combinationSum(tt.candidates, tt.target)

			// 标准化结果和期望输出以便比较
			normalize(result)
			normalize(tt.expected)

			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("combinationSum(%v, %d) = %v, want %v",
					tt.candidates, tt.target, result, tt.expected)
			}
		})
	}
}

// 测试前需要对结果进行标准化处理，因为题目允许以任意顺序返回组合
func normalize(result [][]int) {
	// 对每个内部切片排序
	for i := range result {
		sort.Ints(result[i])
	}
	// 对外部切片排序
	sort.Slice(result, func(i, j int) bool {
		for k := 0; k < len(result[i]) && k < len(result[j]); k++ {
			if result[i][k] != result[j][k] {
				return result[i][k] < result[j][k]
			}
		}
		return len(result[i]) < len(result[j])
	})
}
