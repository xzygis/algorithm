package _79_word_search

import "testing"

func TestExist(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		word     string
		expected bool
	}{
		{
			name: "Example1",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCCED",
			expected: true,
		},
		{
			name: "Example2",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "SEE",
			expected: true,
		},
		{
			name: "Example3",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCB",
			expected: false,
		},
		{
			name: "SingleChar",
			board: [][]byte{
				{'A'},
			},
			word:     "A",
			expected: true,
		},
		{
			name: "SingleCharNoMatch",
			board: [][]byte{
				{'A'},
			},
			word:     "B",
			expected: false,
		},
		{
			name: "EmptyWord",
			board: [][]byte{
				{'A', 'B'},
				{'C', 'D'},
			},
			word:     "",
			expected: true, // 空字符串应该返回true
		},
		{
			name: "1x1Grid",
			board: [][]byte{
				{'X'},
			},
			word:     "X",
			expected: true,
		},
		{
			name: "1x1GridNoMatch",
			board: [][]byte{
				{'X'},
			},
			word:     "Y",
			expected: false,
		},
		{
			name: "ReusePrevention",
			board: [][]byte{
				{'A', 'A'},
				{'A', 'A'},
			},
			word:     "AAAAA", // 需要5个A，但只有4个单元格
			expected: false,
		},
		{
			name: "AllSameChar",
			board: [][]byte{
				{'A', 'A'},
				{'A', 'A'},
			},
			word:     "AAA",
			expected: true,
		},
		{
			name: "WordLongerThanGrid",
			board: [][]byte{
				{'A'},
			},
			word:     "AAA",
			expected: false,
		},
		{
			name: "CircularPath",
			board: [][]byte{
				{'A', 'B'},
				{'D', 'C'},
			},
			word:     "ABCD", // 需要形成环但无法回溯
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := exist(tt.board, tt.word)
			if result != tt.expected {
				t.Errorf("exist() = %v, want %v for test %s", result, tt.expected, tt.name)
			}
		})
	}
}
