package _34_find_first_and_last_position_of_element_in_sorted_array

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		// 题目提供的示例
		{
			name:     "Example 1: multiple targets",
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   8,
			expected: []int{3, 4},
		},
		{
			name:     "Example 2: target not found",
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   6,
			expected: []int{-1, -1},
		},
		{
			name:     "Example 3: empty array",
			nums:     []int{},
			target:   0,
			expected: []int{-1, -1},
		},
		// 常见边界情况
		{
			name:     "Single occurrence",
			nums:     []int{1, 2, 3, 4, 5},
			target:   3,
			expected: []int{2, 2},
		},
		{
			name:     "Target at beginning",
			nums:     []int{1, 1, 2, 3, 4},
			target:   1,
			expected: []int{0, 1},
		},
		{
			name:     "Target at end",
			nums:     []int{1, 2, 3, 4, 4},
			target:   4,
			expected: []int{3, 4},
		},
		{
			name:     "All elements are target",
			nums:     []int{2, 2, 2, 2},
			target:   2,
			expected: []int{0, 3},
		},
		{
			name:     "Target smaller than all elements",
			nums:     []int{1, 2, 3, 4},
			target:   0,
			expected: []int{-1, -1},
		},
		{
			name:     "Target larger than all elements",
			nums:     []int{1, 2, 3, 4},
			target:   5,
			expected: []int{-1, -1},
		},
		// 特殊案例：数组中只有一个元素
		{
			name:     "Single element array - target exists",
			nums:     []int{5},
			target:   5,
			expected: []int{0, 0},
		},
		{
			name:     "Single element array - target not found",
			nums:     []int{5},
			target:   3,
			expected: []int{-1, -1},
		},
		// 目标值不在数组中，但位于数组元素范围内
		{
			name:     "Target within range but not found",
			nums:     []int{1, 3, 5, 7},
			target:   4,
			expected: []int{-1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := searchRange(tt.nums, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("searchRange(%v, %d) = %v, want %v", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
