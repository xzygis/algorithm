package _3_longest_substring_without_repeating_characters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "示例1",
			s:        "abcabcbb",
			expected: 3, // "abc"
		},
		{
			name:     "示例2",
			s:        "bbbbb",
			expected: 1, // "b"
		},
		{
			name:     "示例3",
			s:        "pwwkew",
			expected: 3, // "wke" 或 "kew"
		},
		{
			name:     "空字符串",
			s:        "",
			expected: 0,
		},
		{
			name:     "单个字符",
			s:        "a",
			expected: 1,
		},
		{
			name:     "全不同字符",
			s:        "abcdefg",
			expected: 7,
		},
		{
			name:     "中间有重复",
			s:        "abccba",
			expected: 3, // "abc" 或 "cba"
		},
		{
			name:     "重复在末尾",
			s:        "abcdeff",
			expected: 6, // "abcdef"
		},
		{
			name:     "重复在开头",
			s:        "aabcdef",
			expected: 6, // "abcdef"
		},
		{
			name:     "复杂情况1",
			s:        "dvdf",
			expected: 3, // "vdf"
		},
		{
			name:     "复杂情况2",
			s:        "anviaj",
			expected: 5, // "nviaj"
		},
		{
			name:     "包含空格",
			s:        "a b c d e",
			expected: 3,
		},
		{
			name:     "包含数字和符号",
			s:        "a1b2c3!@#",
			expected: 9, // 整个字符串
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lengthOfLongestSubstring(tt.s)
			if result != tt.expected {
				t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d",
					tt.s, result, tt.expected)
			}
		})
	}
}
