package _394_decode_string

import "testing"

func Test_DecodeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Example 1", "3[a]2[bc]", "aaabcbc"},
		{"Example 2", "3[a2[c]]", "accaccacc"},
		{"Example 3", "2[abc]3[cd]ef", "abcabccdcdcdef"},
		{"Example 4", "abc3[cd]xyz", "abccdcdcdxyz"},
		{"Single letter", "a", "a"},
		{"Empty brackets", "2[]", ""},
		{"Nested empty", "3[a2[]b]", "ababab"},
		{"Multiple digits", "10[a]", "aaaaaaaaaa"},
		{"Deep nesting", "3[a2[b3[c]]]", "abcccbcccabcccbcccabcccbccc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.input); got != tt.expected {
				t.Errorf("decodeString(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
