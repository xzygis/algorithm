package _62_unique_paths

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name     string
		m        int
		n        int
		expected int
	}{
		{
			name:     "示例1: 3x7网格",
			m:        3,
			n:        7,
			expected: 28,
		},
		{
			name:     "示例2: 3x2网格",
			m:        3,
			n:        2,
			expected: 3,
		},
		{
			name:     "示例3: 7x3网格",
			m:        7,
			n:        3,
			expected: 28,
		},
		{
			name:     "示例4: 3x3网格",
			m:        3,
			n:        3,
			expected: 6,
		},
		{
			name:     "1x1网格",
			m:        1,
			n:        1,
			expected: 1,
		},
		{
			name:     "1x5网格（单行）",
			m:        1,
			n:        5,
			expected: 1,
		},
		{
			name:     "5x1网格（单列）",
			m:        5,
			n:        1,
			expected: 1,
		},
		{
			name:     "2x2网格",
			m:        2,
			n:        2,
			expected: 2,
		},
		{
			name:     "边界值测试-最小网格",
			m:        1,
			n:        1,
			expected: 1,
		},
		{
			name:     "中等规模测试",
			m:        10,
			n:        10,
			expected: 48620,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := uniquePaths(tt.m, tt.n)
			if result != tt.expected {
				t.Errorf("uniquePaths(%d, %d) = %d, want %d",
					tt.m, tt.n, result, tt.expected)
			}
		})
	}
}
