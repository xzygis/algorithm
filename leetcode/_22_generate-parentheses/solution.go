package _22_generate_parentheses

func generateParenthesis(n int) []string {
	var path string
	var ans []string
	backtrace(n, 0, 0, path, &ans)
	return ans
}

func backtrace(n int, open int, close int, path string, ans *[]string) {
	if len(path) == 2*n {
		*ans = append(*ans, path)
		return
	}

	if open < n {
		path += "("
		backtrace(n, open+1, close, path, ans)
		path = path[0 : len(path)-1]
	}

	if close < open {
		path += ")"
		backtrace(n, open, close+1, path, ans)
		path = path[0 : len(path)-1]
	}
}

func generateParenthesisV1(n int) []string {
	var path string
	var ans []string
	generate(n, 0, path, &ans)
	return ans
}

func generate(n int, start int, path string, ans *[]string) {
	if start == 2*n {
		if valid(path) {
			*ans = append(*ans, path)
		}
		return
	}

	path += "("
	generate(n, start+1, path, ans)
	path = path[0 : len(path)-1]
	path += ")"
	generate(n, start+1, path, ans)
}

func valid(path string) bool {
	val := 0
	for _, ch := range path {
		if ch == '(' {
			val++
		} else {
			val--
		}

		if val < 0 {
			return false
		}
	}

	return val == 0
}
