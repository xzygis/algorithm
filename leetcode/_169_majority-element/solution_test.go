package _169_majority_element

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "示例1",
			nums:     []int{3, 2, 3},
			expected: 3,
		},
		{
			name:     "示例2",
			nums:     []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
		{
			name:     "单个元素",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "全相同元素",
			nums:     []int{7, 7, 7, 7, 7},
			expected: 7,
		},
		{
			name:     "负数测试",
			nums:     []int{-1, -1, 2, -1, 2, -1, -1},
			expected: -1,
		},
		{
			name:     "边界值: 最大值",
			nums:     []int{1000000000, 1000000000, 1},
			expected: 1000000000,
		},
		{
			name:     "边界值: 最小值",
			nums:     []int{-1000000000, -1000000000, 1},
			expected: -1000000000,
		},
		{
			name:     "交替出现",
			nums:     []int{1, 2, 1, 2, 1},
			expected: 1,
		},
		{
			name:     "多数元素在开头",
			nums:     []int{5, 5, 5, 1, 2, 3},
			expected: 5,
		},
		{
			name:     "多数元素在结尾",
			nums:     []int{1, 2, 3, 4, 4, 4, 4},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := majorityElement(tt.nums)
			if result != tt.expected {
				t.Errorf("majorityElement(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
