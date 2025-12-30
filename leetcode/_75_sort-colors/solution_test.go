package _75_sort_colors

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "示例1",
			nums:     []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "示例2",
			nums:     []int{2, 0, 1},
			expected: []int{0, 1, 2},
		},
		{
			name:     "全0数组",
			nums:     []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "全1数组",
			nums:     []int{1, 1, 1},
			expected: []int{1, 1, 1},
		},
		{
			name:     "全2数组",
			nums:     []int{2, 2, 2},
			expected: []int{2, 2, 2},
		},
		{
			name:     "已排序数组",
			nums:     []int{0, 0, 1, 1, 2, 2},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "逆序数组",
			nums:     []int{2, 2, 1, 1, 0, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "单元素0",
			nums:     []int{0},
			expected: []int{0},
		},
		{
			name:     "单元素1",
			nums:     []int{1},
			expected: []int{1},
		},
		{
			name:     "单元素2",
			nums:     []int{2},
			expected: []int{2},
		},
		{
			name:     "混合情况1",
			nums:     []int{1, 0, 2},
			expected: []int{0, 1, 2},
		},
		{
			name:     "混合情况2",
			nums:     []int{2, 1, 0, 1, 2, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "边界最大值",
			nums:     make([]int, 300), // 最大长度300
			expected: make([]int, 300),
		},
		{
			name:     "只有0和1",
			nums:     []int{1, 0, 1, 0, 1},
			expected: []int{0, 0, 1, 1, 1},
		},
		{
			name:     "只有0和2",
			nums:     []int{2, 0, 2, 0, 2},
			expected: []int{0, 0, 2, 2, 2},
		},
		{
			name:     "只有1和2",
			nums:     []int{2, 1, 2, 1, 2},
			expected: []int{1, 1, 2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制原数组避免修改影响测试
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			sortColors(numsCopy)

			if !reflect.DeepEqual(numsCopy, tt.expected) {
				t.Errorf("sortColors(%v) = %v, want %v", tt.nums, numsCopy, tt.expected)
			}
		})
	}
}
