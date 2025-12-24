package _1143_longest_common_subsequence

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		text1    string
		text2    string
		expected int
	}{
		// LeetCode 官方示例
		{
			name:     "Example 1",
			text1:    "abcde",
			text2:    "ace",
			expected: 3, // LCS 是 "ace"
		},
		{
			name:     "Example 2",
			text1:    "abc",
			text2:    "abc",
			expected: 3, // LCS 是 "abc"
		},
		{
			name:     "Example 3",
			text1:    "abc",
			text2:    "def",
			expected: 0, // 没有公共子序列
		},
		// 边界情况测试
		{
			name:     "First string empty",
			text1:    "",
			text2:    "abc",
			expected: 0,
		},
		{
			name:     "Second string empty",
			text1:    "abc",
			text2:    "",
			expected: 0,
		},
		{
			name:     "Both strings empty",
			text1:    "",
			text2:    "",
			expected: 0,
		},
		{
			name:     "Single char match",
			text1:    "a",
			text2:    "a",
			expected: 1,
		},
		{
			name:     "Single char no match",
			text1:    "a",
			text2:    "b",
			expected: 0,
		},
		// 复杂情况测试
		{
			name:     "Subsequence not contiguous",
			text1:    "abcde",
			text2:    "ace",
			expected: 3, // LCS 是 "ace"
		},
		{
			name:     "Common subsequence with gaps",
			text1:    "bsbininm",
			text2:    "jmjkbkjkv",
			expected: 1, // LCS 可能包含 'b' 或 'k'
		},
		{
			name:     "All same characters",
			text1:    "aaa",
			text2:    "aaa",
			expected: 3,
		},
		{
			name:     "One string is subsequence of other",
			text1:    "abcdefg",
			text2:    "aceg",
			expected: 4, // LCS 是 "aceg"
		},
		{
			name:     "Repeated characters",
			text1:    "aabbcc",
			text2:    "abc",
			expected: 3, // LCS 是 "abc"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestCommonSubsequence(tt.text1, tt.text2)
			if result != tt.expected {
				t.Errorf("longestCommonSubsequence(%q, %q) = %d, want %d", tt.text1, tt.text2, result, tt.expected)
			}
		})
	}
}
