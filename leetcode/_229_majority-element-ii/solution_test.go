package _229_majority_element_ii

import (
	"sort"
	"testing"
)

// 辅助函数：比较两个切片是否相等（忽略顺序）
func equalSliceIgnoreOrder(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestMajorityElementII(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "负数测试",
			nums:     []int{-1, -1, -1, 2, 2, 2, 3},
			expected: []int{-1, 2},
		},
		{
			name:     "示例1",
			nums:     []int{3, 2, 3},
			expected: []int{3},
		},
		{
			name:     "示例2",
			nums:     []int{1},
			expected: []int{1},
		},
		{
			name:     "示例3",
			nums:     []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "两个多数元素",
			nums:     []int{1, 1, 1, 3, 3, 2, 2, 2},
			expected: []int{1, 2},
		},
		{
			name:     "只有一个多数元素",
			nums:     []int{0, 0, 0, 1, 1, 2, 2},
			expected: []int{0},
		},
		{
			name:     "没有多数元素",
			nums:     []int{1, 2, 3, 4},
			expected: []int{},
		},
		{
			name:     "全相同元素",
			nums:     []int{5, 5, 5, 5, 5},
			expected: []int{5},
		},
		{
			name:     "边界值: 最大数组长度",
			nums:     make([]int, 50000),
			expected: []int{0}, // 全0数组
		},
		{
			name:     "三个候选数但只有两个满足条件",
			nums:     []int{1, 1, 1, 2, 2, 2, 3, 3},
			expected: []int{1, 2},
		},
		{
			name:     "恰好超过n/3",
			nums:     []int{1, 1, 1, 2, 3, 4, 5, 6, 7}, // 9个元素，n/3=3，1出现3次刚好超过
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := majorityElement(tt.nums)
			if !equalSliceIgnoreOrder(result, tt.expected) {
				t.Errorf("majorityElement(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}

// 验证算法正确性
func TestMajorityElementIICorrectness(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected []int
		desc     string
	}{
		{[]int{1, 2, 3, 1, 2, 3, 1, 2}, []int{1, 2}, "两个多数元素交替"},
		{[]int{4, 2, 1, 1}, []int{1}, "只有一个多数元素"},
		{[]int{6, 5, 5}, []int{5}, "简单多数"},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := majorityElement(tc.nums)
			if !equalSliceIgnoreOrder(result, tc.expected) {
				t.Errorf("%s: got %v, want %v", tc.desc, result, tc.expected)
			}
		})
	}
}
