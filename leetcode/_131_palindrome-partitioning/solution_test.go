package _131_palindrome_partitioning

import (
	"reflect"
	"sort"
	"testing"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected [][]string
	}{
		{
			name:     "Example1",
			s:        "aab",
			expected: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			name:     "Example2",
			s:        "a",
			expected: [][]string{{"a"}},
		},
		{
			name:     "EmptyString",
			s:        "",
			expected: [][]string{{}}, // 或 [][]string{} 取决于实现
		},
		{
			name:     "SingleChar",
			s:        "b",
			expected: [][]string{{"b"}},
		},
		{
			name: "AllSame",
			s:    "aaa",
			expected: [][]string{
				{"a", "a", "a"},
				{"a", "aa"},
				{"aa", "a"},
				{"aaa"},
			},
		},
		{
			name:     "NotPalindrome",
			s:        "abc",
			expected: [][]string{{"a", "b", "c"}},
		},
		{
			name: "Racecar",
			s:    "racecar",
			expected: [][]string{
				{"r", "a", "c", "e", "c", "a", "r"},
				{"r", "a", "cec", "a", "r"},
				{"r", "aceca", "r"},
				{"racecar"},
			},
		},
		{
			name: "Abba",
			s:    "abba",
			expected: [][]string{
				{"a", "b", "b", "a"},
				{"a", "bb", "a"},
				{"abba"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := partition(tt.s)

			// 标准化结果和期望输出以便比较
			normalizePartitionResult(result)
			normalizePartitionResult(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("partition(%q) = %v, want %v", tt.s, result, tt.expected)
			}
		})
	}
}

// 辅助函数：对结果进行标准化排序以便比较
// 由于回溯算法产生的顺序可能不同，但内容应该相同，需要排序后比较
func normalizePartitionResult(result [][]string) {
	// 对每个内部切片排序
	for i := range result {
		sort.Strings(result[i])
	}
	// 对外部切片排序
	sort.Slice(result, func(i, j int) bool {
		// 先按长度排序
		if len(result[i]) != len(result[j]) {
			return len(result[i]) < len(result[j])
		}
		// 长度相同则按字典序比较每个元素
		for k := 0; k < len(result[i]); k++ {
			if result[i][k] != result[j][k] {
				return result[i][k] < result[j][k]
			}
		}
		return false
	})
}
