package _198_house_robbe

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"Example 1", []int{1, 2, 3, 1}, 4},
		{"Example 2", []int{2, 7, 9, 3, 1}, 12},
		{"Empty array", []int{}, 0},
		{"Single element", []int{5}, 5},
		{"Two elements - first larger", []int{3, 2}, 3},
		{"Two elements - second larger", []int{1, 2}, 2},
		{"Three elements", []int{1, 2, 3}, 4},
		{"All zeros", []int{0, 0, 0, 0}, 0},
		{"Increasing sequence", []int{1, 3, 2, 4, 5}, 8},
		{"Decreasing sequence", []int{5, 4, 3, 2, 1}, 9},
		{"Large numbers", []int{100, 200, 150, 300}, 500},
		{"Alternating high low", []int{10, 1, 10, 1, 10}, 30},
		{"One large in middle", []int{1, 10, 1, 1, 1}, 11},
		{"All same", []int{5, 5, 5, 5}, 10},
		{"Skip two houses", []int{1, 2, 3, 4, 5, 6}, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rob(tt.nums)
			if result != tt.expected {
				t.Errorf("rob(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
