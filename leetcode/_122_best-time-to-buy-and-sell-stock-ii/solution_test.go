package _122_best_time_to_buy_and_sell_stock_ii

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
			expected: 7,
		},
		{
			name:     "Example 2",
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "Example 3",
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
			name:     "Volatile prices",
			prices:   []int{1, 3, 2, 8, 4, 9},
			expected: 13, // (3-1) + (8-2) + (9-4) = 2 + 6 + 5 = 13
		},
		{
			name:     "Peak at beginning",
			prices:   []int{5, 3, 4, 2, 6},
			expected: 5, // (4-3) + (6-2) = 1 + 4 = 5
		},
		{
			name:     "Large numbers",
			prices:   []int{100, 150, 120, 180, 90, 200},
			expected: 220, // (150-100) + (180-120) + (200-90) = 50 + 60 + 110 = 220
		},
		{
			name:     "Gradual increase",
			prices:   []int{1, 2, 3, 4, 5, 6},
			expected: 5, // 6-1 = 5
		},
		{
			name:     "Gradual decrease",
			prices:   []int{6, 5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			name:     "Multiple peaks and valleys",
			prices:   []int{2, 4, 1, 7, 5, 9, 3, 6},
			expected: 15, // (4-2) + (7-1) + (9-5) + (6-3) = 2 + 6 + 4 + 3 = 15
		},
		{
			name:     "Zero prices",
			prices:   []int{0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "Single transaction best",
			prices:   []int{3, 2, 6, 5, 0, 3},
			expected: 7, // (6-2) + (3-0) = 4 + 3 = 7
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
