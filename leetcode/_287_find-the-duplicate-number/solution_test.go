package _287_find_the_duplicate_number

import "testing"

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "示例1",
			nums:     []int{1, 3, 4, 2, 2},
			expected: 2,
		},
		{
			name:     "示例2",
			nums:     []int{3, 1, 3, 4, 2},
			expected: 3,
		},
		{
			name:     "示例3: 全相同",
			nums:     []int{3, 3, 3, 3, 3},
			expected: 3,
		},
		{
			name:     "重复在开头",
			nums:     []int{2, 2, 1, 3, 4},
			expected: 2,
		},
		{
			name:     "重复在结尾",
			nums:     []int{1, 2, 3, 4, 4},
			expected: 4,
		},
		{
			name:     "小数组",
			nums:     []int{1, 1},
			expected: 1,
		},
		{
			name:     "三个元素",
			nums:     []int{2, 2, 2},
			expected: 2,
		},
		{
			name:     "多个重复但只有一个数字",
			nums:     []int{5, 1, 2, 3, 4, 5, 5},
			expected: 5,
		},
		{
			name:     "边界测试: 最小值",
			nums:     []int{1, 1, 2, 3, 4},
			expected: 1,
		},
		{
			name:     "边界测试: 最大值",
			nums:     []int{1, 2, 3, 4, 5, 5},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findDuplicate(tt.nums)
			if result != tt.expected {
				t.Errorf("findDuplicate(%v) = %d, want %d",
					tt.nums, result, tt.expected)
			}
		})
	}
}
