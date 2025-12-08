package _139_word_break

import "testing"

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		expected bool
	}{
		{
			name:     "Example 1",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			expected: true,
		},
		{
			name:     "Example 2",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			expected: true,
		},
		{
			name:     "Example 3",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expected: false,
		},
		{
			name:     "Single word match",
			s:        "apple",
			wordDict: []string{"apple"},
			expected: true,
		},
		{
			name:     "Single word no match",
			s:        "cat",
			wordDict: []string{"dog"},
			expected: false,
		},
		{
			name:     "Repeated words",
			s:        "catcat",
			wordDict: []string{"cat"},
			expected: true,
		},
		{
			name:     "Multiple splits needed",
			s:        "abcd",
			wordDict: []string{"a", "abc", "d", "cd"},
			expected: true, // a-bcd或abc-d
		},
		{
			name:     "Long sequence",
			s:        "aaaaab",
			wordDict: []string{"a", "aa", "aaa", "aaaa"},
			expected: false, // 无法凑出b
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := wordBreak(tt.s, tt.wordDict)
			if result != tt.expected {
				t.Errorf("wordBreak(%q, %v) = %v, want %v",
					tt.s, tt.wordDict, result, tt.expected)
			}
		})
	}
}
