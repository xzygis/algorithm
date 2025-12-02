package _118_pascals_triangle

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name     string
		numRows  int
		expected [][]int
	}{
		{
			name:     "numRows = 1",
			numRows:  1,
			expected: [][]int{{1}},
		},
		{
			name:     "numRows = 2",
			numRows:  2,
			expected: [][]int{{1}, {1, 1}},
		},
		{
			name:     "numRows = 3",
			numRows:  3,
			expected: [][]int{{1}, {1, 1}, {1, 2, 1}},
		},
		{
			name:     "numRows = 4",
			numRows:  4,
			expected: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}},
		},
		{
			name:    "numRows = 5",
			numRows: 5,
			expected: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
		{
			name:    "numRows = 6",
			numRows: 6,
			expected: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
				{1, 5, 10, 10, 5, 1},
			},
		},
		{
			name:    "numRows = 7",
			numRows: 7,
			expected: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
				{1, 5, 10, 10, 5, 1},
				{1, 6, 15, 20, 15, 6, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generate(tt.numRows)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("generate(%d) = %v, want %v", tt.numRows, result, tt.expected)
			}
		})
	}
}
