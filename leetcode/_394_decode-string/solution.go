package _394_decode_string

import "strconv"

func decodeString(s string) string {
	var stack []string
	ptr := 0
	for ptr < len(s) {
		if s[ptr] >= '0' && s[ptr] <= '9' {
			//收集连续数字
			digits := getDigits(s, &ptr)
			stack = append(stack, digits)
		} else if (s[ptr] >= 'a' && s[ptr] <= 'z') || (s[ptr] >= 'A' && s[ptr] <= 'Z') || s[ptr] == '[' {
			//只要不是右括号，通通入栈
			stack = append(stack, string(s[ptr]))
			ptr++
		} else {
			var words []string
			for stack[len(stack)-1] != "[" {
				words = append([]string{stack[len(stack)-1]}, words...)
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1] // '['出栈
			timesStr := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // 数字出栈
			times, _ := strconv.Atoi(timesStr)
			stack = append(stack, getString(words, times)) // 临时计算结果入栈
			ptr++
		}
	}

	return getString(stack, 1)
}

func getString(words []string, times int) string {
	var ans string
	for i := 0; i < times; i++ {
		for _, word := range words {
			ans += word
		}
	}

	return ans
}

func getDigits(s string, ptr *int) string {
	var ans string
	for s[*ptr] >= '0' && s[*ptr] <= '9' {
		ans += string(s[*ptr])
		*ptr++
	}

	return ans
}
