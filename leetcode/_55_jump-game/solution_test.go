package _55_jump_game

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
		{
			name:     "Single element",
			nums:     []int{0},
			expected: true,
		},
		{
			name:     "Single element non-zero",
			nums:     []int{5},
			expected: true,
		},
		{
			name:     "All zeros except last",
			nums:     []int{0, 0, 0, 1},
			expected: false,
		},
		{
			name:     "Can jump over zero",
			nums:     []int{2, 0, 0},
			expected: true,
		},
		{
			name:     "Cannot jump over zero",
			nums:     []int{1, 0, 2},
			expected: false,
		},
		{
			name:     "Large numbers",
			nums:     []int{5, 0, 0, 0, 0, 0},
			expected: true,
		},
		{
			name:     "Exactly reach end",
			nums:     []int{1, 1, 1, 1},
			expected: true,
		},
		{
			name:     "Cannot reach end by 1",
			nums:     []int{1, 1, 0, 1},
			expected: false,
		},
		{
			name:     "Zero at start",
			nums:     []int{0, 2, 3},
			expected: false,
		},
		{
			name:     "Large jump at start",
			nums:     []int{10, 0, 0, 0, 0, 0},
			expected: true,
		},
		{
			name:     "Multiple zeros in middle",
			nums:     []int{2, 0, 0, 1, 0, 1},
			expected: false,
		},
		{
			name:     "Jump exactly to end",
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
		{
			name:     "All ones",
			nums:     []int{1, 1, 1, 1, 1},
			expected: true,
		},
		{
			name:     "Decreasing sequence",
			nums:     []int{4, 3, 2, 1, 0},
			expected: true,
		},
		{
			name:     "Large array with zero gap",
			nums:     []int{3, 0, 0, 0, 2, 0, 1},
			expected: false,
		},
		{
			name:     "Edge case: zero at end",
			nums:     []int{2, 0, 0},
			expected: true,
		},
		{
			name:     "Complex case 1",
			nums:     []int{2, 5, 0, 0},
			expected: true,
		},
		{
			name:     "Complex case 2",
			nums:     []int{1, 2, 0, 1},
			expected: true,
		},
		{
			name:     "Complex case 3",
			nums:     []int{3, 0, 8, 2, 0, 0, 1},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canJump(tt.nums)
			if result != tt.expected {
				t.Errorf("canJump(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
