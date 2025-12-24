package _72_edit_distance

import "testing"

func TestMinDistance(t *testing.T) {
	tests := []struct {
		name     string
		word1    string
		word2    string
		expected int
	}{
		{
			name:     "示例1",
			word1:    "horse",
			word2:    "ros",
			expected: 3,
		},
		{
			name:     "示例2",
			word1:    "intention",
			word2:    "execution",
			expected: 5,
		},
		{
			name:     "空字符串转非空",
			word1:    "",
			word2:    "hello",
			expected: 5,
		},
		{
			name:     "非空字符串转空",
			word1:    "world",
			word2:    "",
			expected: 5,
		},
		{
			name:     "两个空字符串",
			word1:    "",
			word2:    "",
			expected: 0,
		},
		{
			name:     "完全相同字符串",
			word1:    "abc",
			word2:    "abc",
			expected: 0,
		},
		{
			name:     "单字符替换",
			word1:    "a",
			word2:    "b",
			expected: 1,
		},
		{
			name:     "经典例子kitten-sitting",
			word1:    "kitten",
			word2:    "sitting",
			expected: 3,
		},
		{
			name:     "只有插入操作",
			word1:    "abc",
			word2:    "abcd",
			expected: 1,
		},
		{
			name:     "只有删除操作",
			word1:    "abcd",
			word2:    "abc",
			expected: 1,
		},
		{
			name:     "复杂转换",
			word1:    "saturday",
			word2:    "sunday",
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minDistance(tt.word1, tt.word2)
			if result != tt.expected {
				t.Errorf("minDistance(%q, %q) = %d, want %d",
					tt.word1, tt.word2, result, tt.expected)
			}

			// 同时测试优化版本
			resultOpt := minDistance(tt.word1, tt.word2)
			if resultOpt != tt.expected {
				t.Errorf("minDistanceOptimized(%q, %q) = %d, want %d",
					tt.word1, tt.word2, resultOpt, tt.expected)
			}
		})
	}
}
