package _239_sliding_window_maximum

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		// 题目示例
		{
			name:     "Example 1",
			nums:     []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:        3,
			expected: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name:     "Example 2: single element",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "k equals array length",
			nums:     []int{1, 3, 2, 4},
			k:        4,
			expected: []int{4},
		},
		{
			name:     "k equals 1",
			nums:     []int{5, 4, 3, 2, 1},
			k:        1,
			expected: []int{5, 4, 3, 2, 1},
		},
		// 递减序列
		{
			name:     "Descending sequence",
			nums:     []int{5, 4, 3, 2, 1},
			k:        3,
			expected: []int{5, 4, 3},
		},
		// 递增序列
		{
			name:     "Ascending sequence",
			nums:     []int{1, 2, 3, 4, 5},
			k:        3,
			expected: []int{3, 4, 5},
		},
		// 相同元素
		{
			name:     "All same elements",
			nums:     []int{2, 2, 2, 2, 2},
			k:        3,
			expected: []int{2, 2, 2},
		},
		// 负数和零
		{
			name:     "With negatives and zero",
			nums:     []int{-1, -3, -2, 0, -5, 4},
			k:        3,
			expected: []int{-1, 0, 0, 4},
		},
		// 大窗口
		{
			name:     "Large k",
			nums:     []int{9, 7, 2, 4, 6, 8, 2, 1, 5},
			k:        5,
			expected: []int{9, 8, 8, 8, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxSlidingWindow(tt.nums, tt.k)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("maxSlidingWindow(%v, %d) = %v, want %v",
					tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}
