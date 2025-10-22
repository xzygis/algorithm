package _739_daily_temperatures

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Example 1", []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{"Example 2", []int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{"Example 3", []int{30, 60, 90}, []int{1, 1, 0}},
		{"Single day", []int{50}, []int{0}},
		{"All same", []int{30, 30, 30}, []int{0, 0, 0}},
		{"Descending", []int{80, 70, 60, 50}, []int{0, 0, 0, 0}},
		{"Ascending", []int{50, 60, 70, 80}, []int{1, 1, 1, 0}},
		{"Mixed", []int{55, 60, 50, 55, 60, 70}, []int{1, 4, 1, 1, 1, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dailyTemperatures(tt.input); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("dailyTemperatures(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
