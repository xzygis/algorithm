package _137_single_number_ii

import "testing"

func TestSingleNumberII(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "示例1",
			nums:     []int{2, 2, 3, 2},
			expected: 3,
		},
		{
			name:     "示例2",
			nums:     []int{0, 1, 0, 1, 0, 1, 99},
			expected: 99,
		},
		{
			name:     "单个元素",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "负数测试",
			nums:     []int{-2, -2, -2, 1},
			expected: 1,
		},
		{
			name:     "多个负数",
			nums:     []int{-3, -3, -3, -4},
			expected: -4,
		},
		{
			name:     "零值测试",
			nums:     []int{0, 0, 0, 7},
			expected: 7,
		},
		{
			name:     "全零但有一个非零",
			nums:     []int{0, 0, 0, 8},
			expected: 8,
		},
		{
			name:     "出现三次的数字在前",
			nums:     []int{1, 1, 1, 2},
			expected: 2,
		},
		{
			name:     "出现三次的数字在后",
			nums:     []int{2, 1, 1, 1},
			expected: 2,
		},
		{
			name:     "边界值: 最大正数",
			nums:     []int{2147483647, 2147483647, 2147483647, 1},
			expected: 1,
		},
		{
			name:     "边界值: 最小负数",
			nums:     []int{-2147483648, -2147483648, -2147483648, 1},
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
