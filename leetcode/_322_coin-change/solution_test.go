package _322_coin_change

import "testing"

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name     string
		coins    []int
		amount   int
		expected int
	}{
		// LeetCode 官方示例
		{
			name:     "Example 1",
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 3, // 5+5+1
		},
		{
			name:     "Example 2",
			coins:    []int{2},
			amount:   3,
			expected: -1, // 无法凑出
		},
		{
			name:     "Example 3",
			coins:    []int{1},
			amount:   0,
			expected: 0, // 金额为0
		},

		// 边界情况测试
		{
			name:     "Single coin exact match",
			coins:    []int{5},
			amount:   5,
			expected: 1, // 刚好一个硬币
		},
		{
			name:     "Single coin multiple",
			coins:    []int{2},
			amount:   6,
			expected: 3, // 2+2+2
		},
		{
			name:     "Amount less than smallest coin",
			coins:    []int{3, 5, 7},
			amount:   2,
			expected: -1, // 无法凑出
		},

		// 多种硬币组合测试
		{
			name:     "Multiple coins optimal solution",
			coins:    []int{1, 3, 4},
			amount:   6,
			expected: 2, // 3+3
		},
		{
			name:     "All coins needed",
			coins:    []int{1, 2},
			amount:   3,
			expected: 2, // 2+1
		},
		{
			name:     "Large amount with small coins",
			coins:    []int{1, 2, 5},
			amount:   100,
			expected: 20, // 20个5元硬币
		},

		// 特殊序列测试
		{
			name:     "Prime coins",
			coins:    []int{2, 3, 5, 7},
			amount:   11,
			expected: 3,
		},
		{
			name:     "Coin value 1 always works",
			coins:    []int{1, 5, 10},
			amount:   27,
			expected: 5, // 10+10+5+1+1
		},

		// 无法凑出的情况
		{
			name:     "No combination possible",
			coins:    []int{3, 7},
			amount:   5,
			expected: -1,
		},
		{
			name:     "Large amount impossible",
			coins:    []int{2, 4, 8},
			amount:   7,
			expected: -1,
		},

		// 空硬币数组测试
		{
			name:     "Empty coins array",
			coins:    []int{},
			amount:   5,
			expected: -1,
		},

		// 大数测试
		{
			name:     "Large coin values",
			coins:    []int{1000, 5000, 10000},
			amount:   7500,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 特殊处理：如果硬币数组为空，测试逻辑
			if len(tt.coins) == 0 && tt.amount == 0 {
				result := coinChange(tt.coins, tt.amount)
				if result != 0 {
					t.Errorf("coinChange(%v, %d) = %d, want %d",
						tt.coins, tt.amount, result, 0)
				}
				return
			}

			result := coinChange(tt.coins, tt.amount)
			if result != tt.expected {
				t.Errorf("coinChange(%v, %d) = %d, want %d",
					tt.coins, tt.amount, result, tt.expected)
			}
		})
	}
}
