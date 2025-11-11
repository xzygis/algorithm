package _215_kth_largest_element_in_an_array

import "testing"

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
		{
			name:     "K=1 (largest)",
			nums:     []int{1, 2, 3, 4, 5},
			k:        1,
			expected: 5,
		},
		{
			name:     "K=length (smallest)",
			nums:     []int{1, 2, 3, 4, 5},
			k:        5,
			expected: 1,
		},
		{
			name:     "All same elements",
			nums:     []int{7, 7, 7, 7, 7},
			k:        3,
			expected: 7,
		},
		{
			name:     "Single element",
			nums:     []int{42},
			k:        1,
			expected: 42,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			k:        2,
			expected: -2,
		},
		{
			name:     "Mixed positive negative",
			nums:     []int{-1, 2, -3, 4, -5},
			k:        3,
			expected: -1,
		},
		{
			name:     "Descending order",
			nums:     []int{5, 4, 3, 2, 1},
			k:        3,
			expected: 3,
		},
		{
			name:     "Ascending order",
			nums:     []int{1, 2, 3, 4, 5},
			k:        3,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制数组，避免修改原数组影响其他测试
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			if got := findKthLargest(numsCopy, tt.k); got != tt.expected {
				t.Errorf("findKthLargest(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.expected)
			}
		})
	}
}
