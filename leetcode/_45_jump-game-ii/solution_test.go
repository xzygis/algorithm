package _45_jump_game_ii

import "testing"

func TestJump(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 3, 1, 1, 4},
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{2, 3, 0, 1, 4},
			expected: 2,
		},
		{
			name:     "Single element",
			nums:     []int{0},
			expected: 0,
		},
		{
			name:     "Two elements",
			nums:     []int{1, 2},
			expected: 1,
		},
		{
			name:     "Already at end",
			nums:     []int{5},
			expected: 0,
		},
		{
			name:     "All ones",
			nums:     []int{1, 1, 1, 1, 1},
			expected: 4,
		},
		{
			name:     "Large jump at start",
			nums:     []int{5, 1, 1, 1, 1, 1},
			expected: 1,
		},
		{
			name:     "Need multiple jumps",
			nums:     []int{1, 2, 1, 1, 1},
			expected: 3,
		},
		{
			name:     "Optimal path not greedy",
			nums:     []int{2, 1, 3, 1, 1, 1},
			expected: 2,
		},
		{
			name:     "Large numbers",
			nums:     []int{10, 1, 1, 1, 1, 1},
			expected: 1,
		},
		{
			name:     "Zero in middle but can skip",
			nums:     []int{3, 0, 0, 2, 0, 1},
			expected: 2,
		},
		{
			name:     "Gradual increase",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 3,
		},
		{
			name:     "Decrease then increase",
			nums:     []int{3, 2, 1, 4, 1},
			expected: 2,
		},
		{
			name:     "Complex case 1",
			nums:     []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0},
			expected: 3,
		},
		{
			name:     "Complex case 2",
			nums:     []int{1, 1, 1, 1, 1, 4},
			expected: 5,
		},
		{
			name:     "Edge case - minimal jumps",
			nums:     []int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3},
			expected: 2,
		},
		{
			name:     "Need to choose optimal path",
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := jump(tt.nums)
			if result != tt.expected {
				t.Errorf("jump(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
