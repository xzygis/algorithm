package _20_valid_parentheses

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		// 题目提供的示例
		{
			name:     "Example 1: simple parentheses",
			s:        "()",
			expected: true,
		},
		{
			name:     "Example 2: multiple bracket types",
			s:        "()[]{}",
			expected: true,
		},
		{
			name:     "Example 3: mismatched brackets",
			s:        "(]",
			expected: false,
		},
		{
			name:     "Example 4: nested brackets",
			s:        "([])",
			expected: true,
		},
		{
			name:     "Example 5: incorrectly nested brackets",
			s:        "([)]",
			expected: false,
		},
		// 边界情况
		{
			name:     "Empty string",
			s:        "",
			expected: true, // 空字符串通常被认为是有效的
		},
		{
			name:     "Single opening bracket",
			s:        "(",
			expected: false,
		},
		{
			name:     "Single closing bracket",
			s:        ")",
			expected: false,
		},
		// 复杂嵌套
		{
			name:     "Complex valid nesting",
			s:        "{[]}",
			expected: true,
		},
		{
			name:     "Complex invalid nesting",
			s:        "{[}]",
			expected: false,
		},
		{
			name:     "Deeply nested valid",
			s:        "({[()]})",
			expected: true,
		},
		{
			name:     "Deeply nested invalid",
			s:        "({[(])})",
			expected: false,
		},
		// 顺序错误
		{
			name:     "Closing before opening",
			s:        "}{",
			expected: false,
		},
		{
			name:     "Only opening brackets",
			s:        "({[",
			expected: false,
		},
		{
			name:     "Only closing brackets",
			s:        ")}]",
			expected: false,
		},
		{
			name:     "Long valid sequence",
			s:        "()((()))[]{}{[()()]}",
			expected: true,
		},
		{
			name:     "Long invalid sequence",
			s:        "()((()))[]{}{[()(])}",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValid(tt.s)
			if result != tt.expected {
				t.Errorf("isValid(%q) = %v, want %v", tt.s, result, tt.expected)
			}
		})
	}
}
