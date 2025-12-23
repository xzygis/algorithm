package _64_minimum_path_sum

import "testing"

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "示例1",
			grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			expected: 7, // 1→3→1→1→1
		},
		{
			name: "示例2",
			grid: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			expected: 12, // 1→2→3→6
		},
		{
			name: "单行网格",
			grid: [][]int{
				{1, 2, 3, 4},
			},
			expected: 10, // 1+2+3+4
		},
		{
			name: "单列网格",
			grid: [][]int{
				{1},
				{2},
				{3},
			},
			expected: 6, // 1+2+3
		},
		{
			name: "1x1网格",
			grid: [][]int{
				{5},
			},
			expected: 5,
		},
		{
			name: "2x2网格",
			grid: [][]int{
				{1, 2},
				{1, 3},
			},
			expected: 5, // 1→1→3
		},
		{
			name: "包含0的网格",
			grid: [][]int{
				{1, 0, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			expected: 3, // 多条路径，最小和为3
		},
		{
			name: "所有元素相同",
			grid: [][]int{
				{2, 2, 2},
				{2, 2, 2},
				{2, 2, 2},
			},
			expected: 10, // 2×5=10
		},
		{
			name: "最小值边界",
			grid: [][]int{
				{0, 0},
				{0, 0},
			},
			expected: 0,
		},
		{
			name: "最大值边界",
			grid: [][]int{
				{200, 200},
				{200, 200},
			},
			expected: 600,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minPathSum(tt.grid)
			if result != tt.expected {
				t.Errorf("minPathSum(%v) = %d, want %d", tt.grid, result, tt.expected)
			}

			// 同时测试优化版本
			resultOpt := minPathSum(tt.grid)
			if resultOpt != tt.expected {
				t.Errorf("minPathSumOptimized(%v) = %d, want %d", tt.grid, resultOpt, tt.expected)
			}
		})
	}
}
