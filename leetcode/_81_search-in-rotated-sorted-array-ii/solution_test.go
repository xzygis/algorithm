package _81_search_in_rotated_sorted_array_ii

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected bool
	}{
		// 题目提供的示例
		{
			name:     "Example 1: target exists",
			nums:     []int{2, 5, 6, 0, 0, 1, 2},
			target:   0,
			expected: true,
		},
		{
			name:     "Example 2: target does not exist",
			nums:     []int{2, 5, 6, 0, 0, 1, 2},
			target:   3,
			expected: false,
		},
		// 边界情况
		{
			name:     "Empty array",
			nums:     []int{},
			target:   0,
			expected: false,
		},
		{
			name:     "Single element - exists",
			nums:     []int{1},
			target:   1,
			expected: true,
		},
		{
			name:     "Single element - does not exist",
			nums:     []int{1},
			target:   0,
			expected: false,
		},
		// 重复元素处理
		{
			name:     "All same elements - exists",
			nums:     []int{1, 1, 1, 1, 1},
			target:   1,
			expected: true,
		},
		{
			name:     "All same elements - does not exist",
			nums:     []int{1, 1, 1, 1, 1},
			target:   2,
			expected: false,
		},
		{
			name:     "Duplicate elements at boundaries - exists",
			nums:     []int{1, 0, 1, 1, 1}, // 旋转后数组，测试重复元素处理
			target:   0,
			expected: true,
		},
		// 无重复元素（类似LeetCode 33题）
		{
			name:     "No duplicates - exists at rotation point",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: true,
		},
		{
			name:     "No duplicates - does not exist",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: false,
		},
		// 目标值在开头或结尾
		{
			name:     "Target at beginning with duplicates",
			nums:     []int{3, 1, 3, 3, 3},
			target:   3,
			expected: true,
		},
		{
			name:     "Target at end with duplicates",
			nums:     []int{3, 3, 3, 1, 3},
			target:   1,
			expected: true,
		},
		// 复杂重复场景
		{
			name:     "Multiple duplicates and rotation - exists",
			nums:     []int{4, 5, 6, 6, 7, 0, 0, 1, 2, 4, 4},
			target:   5,
			expected: true,
		},
		{
			name:     "Multiple duplicates and rotation - does not exist",
			nums:     []int{4, 5, 6, 6, 7, 0, 0, 1, 2, 4, 4},
			target:   3,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := search(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("search(%v, %d) = %v, want %v", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
