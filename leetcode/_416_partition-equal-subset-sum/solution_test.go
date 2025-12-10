package _416_partition_equal_subset_sum

import "testing"

func TestCanPartition(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "示例1: 可分割情况",
			nums:     []int{1, 5, 11, 5},
			expected: true, // [1,5,5] 和 [11]
		},
		{
			name:     "示例2: 不可分割情况",
			nums:     []int{1, 2, 3, 5},
			expected: false,
		},
		{
			name:     "空数组",
			nums:     []int{},
			expected: false,
		},
		{
			name:     "单个元素",
			nums:     []int{5},
			expected: false, // 无法分成两个非空子集
		},
		{
			name:     "两个相同元素",
			nums:     []int{1, 1},
			expected: true, // [1] 和 [1]
		},
		{
			name:     "两个不同元素",
			nums:     []int{1, 2},
			expected: false, // 1+2=3是奇数
		},
		{
			name:     "总和为奇数",
			nums:     []int{1, 2, 3, 4, 5}, // 和=15
			expected: false,
		},
		{
			name:     "包含重复数字",
			nums:     []int{1, 1, 1, 1},
			expected: true, // [1,1] 和 [1,1]
		},
		{
			name:     "最大值等于target",
			nums:     []int{3, 3, 6},
			expected: true, // [6] 和 [3,3]
		},
		{
			name:     "最大值大于target",
			nums:     []int{2, 2, 5}, // 和=9, target=4.5
			expected: false,
		},
		{
			name:     "复杂可分割案例",
			nums:     []int{1, 2, 3, 4, 5, 6, 7}, // 和=28, target=14
			expected: true,                       // [1,4,6,3] 和 [2,5,7] 等多种分法
		},
		{
			name:     "边界值测试",
			nums:     []int{100, 100}, // 最大允许值
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canPartition(tt.nums)
			if result != tt.expected {
				t.Errorf("canPartition(%v) = %v, want %v",
					tt.nums, result, tt.expected)
			}
		})
	}
}
