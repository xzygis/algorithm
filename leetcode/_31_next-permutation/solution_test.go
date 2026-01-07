package _31_next_permutation

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "示例1: 标准情况",
			nums:     []int{1, 2, 3},
			expected: []int{1, 3, 2},
		},
		{
			name:     "示例2: 最大排列变最小",
			nums:     []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
		{
			name:     "示例3: 有重复元素",
			nums:     []int{1, 1, 5},
			expected: []int{1, 5, 1},
		},
		{
			name:     "单元素数组",
			nums:     []int{1},
			expected: []int{1},
		},
		{
			name:     "双元素数组升序",
			nums:     []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "双元素数组降序",
			nums:     []int{2, 1},
			expected: []int{1, 2},
		},
		{
			name:     "复杂情况1",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 5, 4},
		},
		{
			name:     "复杂情况2",
			nums:     []int{1, 3, 5, 4, 2},
			expected: []int{1, 4, 2, 3, 5},
		},
		{
			name:     "全部相同元素",
			nums:     []int{2, 2, 2, 2},
			expected: []int{2, 2, 2, 2},
		},
		{
			name:     "边界情况: 末尾有降序",
			nums:     []int{1, 2, 3, 2, 1},
			expected: []int{1, 3, 1, 2, 2},
		},
		{
			name:     "中间有跳跃",
			nums:     []int{2, 3, 1},
			expected: []int{3, 1, 2},
		},
		{
			name:     "测试用例1",
			nums:     []int{1, 2, 3, 4},
			expected: []int{1, 2, 4, 3},
		},
		{
			name:     "测试用例2",
			nums:     []int{4, 3, 2, 1},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "测试用例3",
			nums:     []int{1, 3, 2},
			expected: []int{2, 1, 3},
		},
		{
			name:     "测试用例4",
			nums:     []int{1, 5, 1},
			expected: []int{5, 1, 1},
		},
		{
			name:     "测试用例5",
			nums:     []int{5, 1, 1},
			expected: []int{1, 1, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制原数组以避免测试间的相互影响
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)

			nextPermutation(numsCopy)

			if !reflect.DeepEqual(numsCopy, tt.expected) {
				t.Errorf("nextPermutation(%v) = %v, want %v",
					tt.nums, numsCopy, tt.expected)
			}
		})
	}
}
