package _347_top_k_frequent_elements

import (
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "Example 2",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "Example 3",
			nums:     []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -1, -2, -2, -2, -3},
			k:        2,
			expected: []int{-2, -1},
		},
		{
			name:     "k equals unique elements",
			nums:     []int{1, 1, 2, 2, 3},
			k:        3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "Large k",
			nums:     []int{1, 2, 3, 1, 2, 1},
			k:        3,
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := topKFrequent(tt.nums, tt.k)

			// 由于结果顺序不重要，我们需要排序后比较
			sort.Ints(result)
			sort.Ints(tt.expected)

			if !equalSlices(result, tt.expected) {
				t.Errorf("topKFrequent(%v, %d) = %v, want %v",
					tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}

// 辅助函数：比较两个切片是否相等
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
