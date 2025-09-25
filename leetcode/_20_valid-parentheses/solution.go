package _20_valid_parentheses

var mapping = map[byte]byte{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	var stack []byte
	for i := range s {
		ch := s[i] //这种写法ch的类型是uint8,等同于byte。如果是for i, ch这种写法，ch的类型是int32，后续不太方便处理
		if _, ok := mapping[ch]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != mapping[ch] {
				return false
			}

			stack = stack[0 : len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}

	return len(stack) == 0
}
