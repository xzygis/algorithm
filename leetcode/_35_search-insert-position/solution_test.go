package _35_search_insert_position

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		// 题目示例
		{
			name:     "Example 1: Target exists in middle",
			nums:     []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			name:     "Example 2: Target does not exist, insert in middle",
			nums:     []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			name:     "Example 3: Target larger than all elements",
			nums:     []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
		// 边界情况
		{
			name:     "Target smaller than all elements",
			nums:     []int{1, 3, 5, 6},
			target:   0,
			expected: 0,
		},
		{
			name:     "Single element array - target exists",
			nums:     []int{5},
			target:   5,
			expected: 0,
		},
		{
			name:     "Single element array - target does not exist (smaller)",
			nums:     []int{5},
			target:   3,
			expected: 0,
		},
		{
			name:     "Single element array - target does not exist (larger)",
			nums:     []int{5},
			target:   7,
			expected: 1,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			target:   5,
			expected: 0,
		},
		{
			name:     "Target at beginning",
			nums:     []int{1, 3, 5, 6},
			target:   1,
			expected: 0,
		},
		{
			name:     "Target at end",
			nums:     []int{1, 3, 5, 6},
			target:   6,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := searchInsert(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("searchInsert(%v, %d) = %d, want %d", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
