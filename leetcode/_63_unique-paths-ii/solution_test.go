package _63_unique_paths_ii

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "示例1: 中间有障碍物",
			grid: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			expected: 2,
		},
		{
			name: "示例2: 简单障碍物",
			grid: [][]int{
				{0, 1},
				{0, 0},
			},
			expected: 1,
		},
		{
			name: "起点有障碍物",
			grid: [][]int{
				{1, 0},
				{0, 0},
			},
			expected: 0,
		},
		{
			name: "终点有障碍物",
			grid: [][]int{
				{0, 0},
				{0, 1},
			},
			expected: 0,
		},
		{
			name: "无障碍物3x3网格",
			grid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			expected: 6,
		},
		{
			name: "第一行有障碍物",
			grid: [][]int{
				{0, 1, 0},
				{0, 0, 0},
			},
			expected: 1,
		},
		{
			name: "第一列有障碍物",
			grid: [][]int{
				{0, 0},
				{1, 0},
				{0, 0},
			},
			expected: 1,
		},
		{
			name: "多条障碍物阻断路径",
			grid: [][]int{
				{0, 0, 0},
				{0, 1, 1},
				{0, 0, 0},
			},
			expected: 1,
		},
		{
			name: "完全被障碍物阻断",
			grid: [][]int{
				{0, 1, 0},
				{1, 0, 1},
				{0, 1, 0},
			},
			expected: 0,
		},
		{
			name: "单行网格无障碍",
			grid: [][]int{
				{0, 0, 0},
			},
			expected: 1,
		},
		{
			name: "单行网格有障碍",
			grid: [][]int{
				{0, 1, 0},
			},
			expected: 0,
		},
		{
			name: "单列网格无障碍",
			grid: [][]int{
				{0},
				{0},
				{0},
			},
			expected: 1,
		},
		{
			name: "单列网格有障碍",
			grid: [][]int{
				{0},
				{1},
				{0},
			},
			expected: 0,
		},
		{
			name: "边界值测试-最小网格",
			grid: [][]int{
				{0},
			},
			expected: 1,
		},
		{
			name: "边界值测试-最小网格有障碍",
			grid: [][]int{
				{1},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := uniquePathsWithObstacles(tt.grid)
			if result != tt.expected {
				t.Errorf("uniquePathsWithObstacles() = %d, want %d", result, tt.expected)
			}
		})
	}
}
