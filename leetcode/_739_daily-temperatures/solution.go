package _739_daily_temperatures

import "math"

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	var stack []int
	for i, temp := range temperatures {
		// 当前元素大于栈顶元素
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			// 找到栈顶元素的解
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			// 栈顶元素出栈
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return ans
}

func dailyTemperaturesV1(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	next := make([]int, 101)
	for i := range next {
		next[i] = math.MaxInt32
	}
	for i := len(temperatures) - 1; i >= 0; i-- {
		smallestWarmerIndex := math.MaxInt32
		for temp := temperatures[i] + 1; temp <= 100; temp++ {
			if smallestWarmerIndex > next[temp] {
				smallestWarmerIndex = next[temp]
			}
		}

		if smallestWarmerIndex != math.MaxInt32 {
			ans[i] = smallestWarmerIndex - i
		}

		next[temperatures[i]] = i
	}

	return ans
}
