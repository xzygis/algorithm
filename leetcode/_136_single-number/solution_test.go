package _136_single_number

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "示例1: 成对数字中间有单数",
			nums:     []int{2, 2, 1},
			expected: 1,
		},
		{
			name:     "示例2: 多个成对数字",
			nums:     []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			name:     "示例3: 单个元素",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "负数测试",
			nums:     []int{-1, -1, -2},
			expected: -2,
		},
		{
			name:     "多个负数",
			nums:     []int{-3, -4, -4, -3, 5},
			expected: 5,
		},
		{
			name:     "零值测试",
			nums:     []int{0, 1, 0},
			expected: 1,
		},
		{
			name:     "全零但有一个非零",
			nums:     []int{0, 0, 7},
			expected: 7,
		},
		{
			name:     "较大数组测试",
			nums:     []int{1, 2, 3, 4, 5, 1, 2, 3, 4},
			expected: 5,
		},
		{
			name:     "成对数字乱序",
			nums:     []int{7, 3, 5, 3, 7},
			expected: 5,
		},
		{
			name:     "边界值: 最大正数",
			nums:     []int{30000, 1, 30000},
			expected: 1,
		},
		{
			name:     "边界值: 最小负数",
			nums:     []int{-30000, 1, -30000},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := singleNumber(tt.nums)
			if result != tt.expected {
				t.Errorf("singleNumber(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
