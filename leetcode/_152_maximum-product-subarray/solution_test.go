package _152_maximum_product_subarray

import "testing"

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "示例1: 正常正数序列",
			nums:     []int{2, 3, -2, 4},
			expected: 6, // 2 * 3=6
		},
		{
			name:     "示例2: 包含0的序列",
			nums:     []int{-2, 0, -1},
			expected: 0, // 0本身或者包含0的子数组
		},
		{
			name:     "全正数序列",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 120, // 全部相乘
		},
		{
			name:     "全负数序列-奇数个",
			nums:     []int{-2, -3, -1},
			expected: 6, // -2*-3=6
		},
		{
			name:     "全负数序列-偶数个",
			nums:     []int{-2, -3, -4, -1},
			expected: 24, // -2*-3*-4*-1=24
		},
		{
			name:     "单个元素",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "包含0和负数",
			nums:     []int{-2, 0, 3, -4},
			expected: 3, // 3或者-4*-3=12，但实际是3
		},
		{
			name:     "负数为边界情况",
			nums:     []int{-2, 3, -4},
			expected: 24, // -2 * 3*-4=24
		},
		{
			name:     "以负数开头",
			nums:     []int{-1, 2, 3},
			expected: 6, // 2 * 3=6
		},
		{
			name:     "多个零分隔",
			nums:     []int{2, 0, 3, 0, 4},
			expected: 4, // 最大的是4
		},
		{
			name:     "大数测试",
			nums:     []int{10, -10, 10, -10},
			expected: 10000, // 10*-10 * 10*-10=10000
		},
		{
			name:     "递减序列",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 120, // 全部相乘
		},
		{
			name:     "先正后负",
			nums:     []int{1, 2, -3, 4, -5},
			expected: 120, // 全部相乘1 * 2*-3 * 4*-5=120
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProduct(tt.nums)
			if result != tt.expected {
				t.Errorf("maxProduct(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
