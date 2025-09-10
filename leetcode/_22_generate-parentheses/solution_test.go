package _22_generate_parentheses

import (
	"reflect"
	"sort"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []string
	}{
		{
			name:     "Example1",
			n:        3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name:     "Example2",
			n:        1,
			expected: []string{"()"},
		},
		{
			name:     "NZero",
			n:        0,
			expected: []string{""}, // 或 []string{} 取决于实现
		},
		{
			name:     "NTwo",
			n:        2,
			expected: []string{"(())", "()()"},
		},
		{
			name: "NFour",
			n:    4,
			expected: []string{
				"(((())))", "((()()))", "((())())", "((()))()", "(()(()))",
				"(()()())", "(()())()", "(())(())", "(())()()", "()((()))",
				"()(()())", "()(())()", "()()(())", "()()()()",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateParenthesis(tt.n)

			// 对结果和期望值进行排序，因为返回顺序可能不同但内容一致
			sort.Strings(result)
			sort.Strings(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("generateParenthesis(%d) = %v, want %v",
					tt.n, result, tt.expected)
			}
		})
	}
}

// 性能测试：测试较大输入时的性能
func BenchmarkGenerateParenthesis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateParenthesis(8)
	}
}
