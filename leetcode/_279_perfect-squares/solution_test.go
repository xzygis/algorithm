package _279_perfect_squares

import "testing"

func TestNumSquares(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		// LeetCode 示例
		{
			name:     "Example 1 - n=12",
			n:        12,
			expected: 3, // 12 = 4 + 4 + 4
		},
		{
			name:     "Example 2 - n=13",
			n:        13,
			expected: 2, // 13 = 4 + 9
		},

		// 边界情况
		{
			name:     "n=0",
			n:        0,
			expected: 0,
		},
		{
			name:     "n=1",
			n:        1,
			expected: 1, // 1 = 1
		},
		{
			name:     "n=2",
			n:        2,
			expected: 2, // 2 = 1 + 1
		},
		{
			name:     "n=3",
			n:        3,
			expected: 3, // 3 = 1 + 1 + 1
		},

		// 完全平方数本身
		{
			name:     "n=4 (perfect square)",
			n:        4,
			expected: 1, // 4 = 4
		},
		{
			name:     "n=9 (perfect square)",
			n:        9,
			expected: 1, // 9 = 9
		},
		{
			name:     "n=16 (perfect square)",
			n:        16,
			expected: 1, // 16 = 16
		},

		// 需要2个平方数的情况
		{
			name:     "n=5 (needs 2 squares)",
			n:        5,
			expected: 2, // 5 = 4 + 1
		},
		{
			name:     "n=8 (needs 2 squares)",
			n:        8,
			expected: 2, // 8 = 4 + 4
		},
		{
			name:     "n=10 (needs 2 squares)",
			n:        10,
			expected: 2, // 10 = 9 + 1
		},

		// 需要3个平方数的情况
		{
			name:     "n=6 (needs 3 squares)",
			n:        6,
			expected: 3, // 6 = 4 + 1 + 1
		},
		{
			name:     "n=7 (needs 3 squares)",
			n:        7,
			expected: 4, // 7 = 4 + 1 + 1 + 1
		},
		{
			name:     "n=11 (needs 3 squares)",
			n:        11,
			expected: 3, // 11 = 9 + 1 + 1
		},

		// 较大数字测试
		{
			name:     "n=25 (perfect square)",
			n:        25,
			expected: 1, // 25 = 25
		},
		{
			name:     "n=26 (needs 2 squares)",
			n:        26,
			expected: 2, // 26 = 25 + 1
		},
		{
			name:     "n=27 (needs 3 squares)",
			n:        27,
			expected: 3, // 27 = 25 + 1 + 1
		},
		{
			name:     "n=28 (needs 4 squares)",
			n:        28,
			expected: 4, // 28 = 25 + 1 + 1 + 1
		},
		{
			name:     "n=29 (needs 2 squares)",
			n:        29,
			expected: 2, // 29 = 25 + 4
		},

		// 需要4个平方数的情况（拉格朗日四平方定理）
		{
			name:     "n=15 (needs 4 squares)",
			n:        15,
			expected: 4, // 15 = 9 + 4 + 1 + 1
		},
		{
			name:     "n=23 (needs 4 squares)",
			n:        23,
			expected: 4, // 23 = 9 + 9 + 4 + 1
		},

		// 最大值边界测试
		{
			name:     "n=10000 (maximum constraint)",
			n:        10000,
			expected: 1, // 10000 = 100^2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numSquares(tt.n)
			if result != tt.expected {
				t.Errorf("numSquares(%d) = %d, want %d", tt.n, result, tt.expected)
			}
		})
	}
}
