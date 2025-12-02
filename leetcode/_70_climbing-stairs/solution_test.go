package _70_climbing_stairs

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "n=0",
			n:        0,
			expected: 0,
		},
		{
			name:     "n=1",
			n:        1,
			expected: 1,
		},
		{
			name:     "n=2",
			n:        2,
			expected: 2,
		},
		{
			name:     "n=3",
			n:        3,
			expected: 3,
		},
		{
			name:     "n=4",
			n:        4,
			expected: 5, // 1+1+1+1, 1+1+2, 1+2+1, 2+1+1, 2+2
		},
		{
			name:     "n=5",
			n:        5,
			expected: 8, // 斐波那契数列
		},
		{
			name:     "n=6",
			n:        6,
			expected: 13,
		},
		{
			name:     "n=7",
			n:        7,
			expected: 21,
		},
		{
			name:     "n=8",
			n:        8,
			expected: 34,
		},
		{
			name:     "n=9",
			n:        9,
			expected: 55,
		},
		{
			name:     "n=10",
			n:        10,
			expected: 89,
		},
		{
			name:     "n=20",
			n:        20,
			expected: 10946,
		},
		{
			name:     "n=30",
			n:        30,
			expected: 1346269,
		},
		{
			name:     "n=40",
			n:        40,
			expected: 165580141,
		},
		{
			name:     "n=45",
			n:        45,
			expected: 1836311903,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := climbStairs(tt.n)
			if result != tt.expected {
				t.Errorf("climbStairs(%d) = %d, want %d", tt.n, result, tt.expected)
			}
		})
	}
}
