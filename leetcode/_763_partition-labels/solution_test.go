package _763_partition_labels

import (
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected []int
	}{
		{
			name:     "Example 1",
			s:        "ababcbacadefegdehijhklij",
			expected: []int{9, 7, 8},
		},
		{
			name:     "Example 2",
			s:        "eccbbbbdec",
			expected: []int{10},
		},
		{
			name:     "Single character",
			s:        "a",
			expected: []int{1},
		},
		{
			name:     "All same characters",
			s:        "aaaaa",
			expected: []int{5},
		},
		{
			name:     "No repetition",
			s:        "abcdefg",
			expected: []int{1, 1, 1, 1, 1, 1, 1},
		},
		{
			name:     "Two segments",
			s:        "abacaba",
			expected: []int{7},
		},
		{
			name:     "Multiple segments",
			s:        "abcadefegd",
			expected: []int{4, 6},
		},
		{
			name:     "Overlapping segments",
			s:        "ababacac",
			expected: []int{8},
		},
		{
			name:     "Complex case 1",
			s:        "qiejxqfnqceocmy",
			expected: []int{13, 1, 1},
		},
		{
			name:     "Complex case 2",
			s:        "caedbdedda",
			expected: []int{1, 9},
		},
		{
			name:     "All characters unique",
			s:        "abcdefghij",
			expected: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name:     "Repeated at beginning and end",
			s:        "abcdea",
			expected: []int{6},
		},
		{
			name:     "Multiple same segments",
			s:        "abccba",
			expected: []int{6},
		},
		{
			name:     "Separate groups",
			s:        "aabbccddeeff",
			expected: []int{2, 2, 2, 2, 2, 2},
		},
		{
			name:     "Interleaved groups",
			s:        "abacbc",
			expected: []int{6},
		},
		{
			name:     "Long repeated pattern",
			s:        "abcabcabc",
			expected: []int{9},
		},
		{
			name:     "Single repeated character",
			s:        "zzzzz",
			expected: []int{5},
		},
		{
			name:     "Alternating pattern",
			s:        "abababab",
			expected: []int{8},
		},
		{
			name:     "Complex overlapping",
			s:        "ababcbacadefegdehijhklijmnopponm",
			expected: []int{9, 7, 8, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := partitionLabels(tt.s)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("partitionLabels(%q) = %v, want %v", tt.s, result, tt.expected)
			}
		})
	}
}
