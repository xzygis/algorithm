package _121_best_time_to_buy_and_sell_stock

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "Example 1",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "Example 2",
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			name:     "Empty array",
			prices:   []int{},
			expected: 0,
		},
		{
			name:     "Single day",
			prices:   []int{5},
			expected: 0,
		},
		{
			name:     "Two days increasing",
			prices:   []int{1, 2},
			expected: 1,
		},
		{
			name:     "Two days decreasing",
			prices:   []int{2, 1},
			expected: 0,
		},
		{
			name:     "All same prices",
			prices:   []int{3, 3, 3, 3, 3},
			expected: 0,
		},
		{
			name:     "Multiple peaks",
			prices:   []int{2, 4, 1, 7, 5, 9},
			expected: 8, // 1买入，9卖出
		},
		{
			name:     "Early peak",
			prices:   []int{3, 2, 6, 5, 0, 3},
			expected: 4, // 2买入，6卖出
		},
		{
			name:     "Large numbers",
			prices:   []int{10000, 9999, 10001, 10002, 9998},
			expected: 3, // 9998买入，10002卖出
		},
		{
			name:     "Zero prices",
			prices:   []int{0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "Gradual increase",
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "Gradual decrease",
			prices:   []int{5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			name:     "Volatile prices",
			prices:   []int{2, 1, 4, 5, 2, 9, 7},
			expected: 8, // 1买入，9卖出
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProfit(tt.prices)
			if result != tt.expected {
				t.Errorf("maxProfit(%v) = %d, want %d", tt.prices, result, tt.expected)
			}
		})
	}
}
