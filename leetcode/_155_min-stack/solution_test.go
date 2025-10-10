package _155_min_stack

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		params     []int
		expected   []interface{}
	}{
		// 题目提供的示例
		{
			name:       "Example 1",
			operations: []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"},
			params:     []int{0, -2, 0, -3, 0, 0, 0, 0},
			expected:   []interface{}{nil, nil, nil, nil, -3, nil, 0, -2},
		},
		// 基本功能测试
		{
			name:       "Basic operations",
			operations: []string{"MinStack", "push", "top", "getMin", "pop", "getMin"},
			params:     []int{0, 5, 0, 0, 0, 0},
			expected:   []interface{}{nil, nil, 5, 5, nil, nil},
		},
		// 最小值变化测试
		{
			name:       "Minimum changes with ascending order",
			operations: []string{"MinStack", "push", "getMin", "push", "getMin", "push", "getMin"},
			params:     []int{0, 3, 0, 2, 0, 1, 0},
			expected:   []interface{}{nil, nil, 3, nil, 2, nil, 1},
		},
		{
			name:       "Minimum changes with descending order",
			operations: []string{"MinStack", "push", "getMin", "push", "getMin", "push", "getMin"},
			params:     []int{0, 1, 0, 2, 0, 3, 0},
			expected:   []interface{}{nil, nil, 1, nil, 1, nil, 1},
		},
		// 重复值测试
		{
			name:       "Duplicate values",
			operations: []string{"MinStack", "push", "push", "getMin", "push", "getMin", "pop", "getMin"},
			params:     []int{0, 2, 2, 0, 2, 0, 0, 0},
			expected:   []interface{}{nil, nil, nil, 2, nil, 2, nil, 2},
		},
		// 边界值测试
		{
			name:       "Boundary values",
			operations: []string{"MinStack", "push", "getMin", "push", "getMin"},
			params:     []int{0, -2147483648, 0, 2147483647, 0},
			expected:   []interface{}{nil, nil, -2147483648, nil, -2147483648},
		},
		// 复杂操作序列
		{
			name:       "Complex operation sequence",
			operations: []string{"MinStack", "push", "push", "push", "pop", "getMin", "push", "getMin", "pop", "top", "getMin"},
			params:     []int{0, 10, 20, 5, 0, 0, 2, 0, 0, 0, 0},
			expected:   []interface{}{nil, nil, nil, nil, nil, 10, nil, 2, nil, 20, 10},
		},
		// 弹出最小值后最小值更新
		{
			name:       "Minimum updates after pop",
			operations: []string{"MinStack", "push", "push", "getMin", "pop", "getMin"},
			params:     []int{0, -3, -2, 0, 0, 0},
			expected:   []interface{}{nil, nil, nil, -3, nil, -2},
		},
		// 单元素栈
		{
			name:       "Single element stack",
			operations: []string{"MinStack", "push", "top", "getMin", "pop"},
			params:     []int{0, 7, 0, 0, 0},
			expected:   []interface{}{nil, nil, 7, 7, nil},
		},
		// 空栈操作（根据题目要求，不会在空栈上调用pop/top/getMin）
		// 但为了健壮性，可以测试初始化后直接getMin的情况（应返回零值或错误，但题目要求非空栈）
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var minStack MinStack

			for i, op := range tt.operations {
				switch op {
				case "MinStack":
					minStack = Constructor()
				case "push":
					minStack.Push(tt.params[i])
				case "pop":
					minStack.Pop()
				case "top":
					result := minStack.Top()
					expected := tt.expected[i].(int)
					if result != expected {
						t.Errorf("Top() at step %d = %d, want %d", i, result, expected)
					}
				case "getMin":
					result := minStack.GetMin()
					expected := tt.expected[i].(int)
					if result != expected {
						t.Errorf("GetMin() at step %d = %d, want %d", i, result, expected)
					}
				}
			}
		})
	}
}
