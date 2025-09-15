package _131_palindrome_partitioning

func partition(s string) [][]string {
	n := len(s)
	f := make([][]bool, n)
	for i := 0; i < n; i++ {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	var path []string
	var ans [][]string
	backtraceV2(s, 0, f, path, &ans)
	return ans
}

func backtraceV2(s string, start int, flag [][]bool, path []string, ans *[][]string) {
	if start == len(s) {
		*ans = append(*ans, append([]string{}, path...))
		return
	}

	// 需要注意：这里需要包含i等于len(s)的情况
	for i := start + 1; i <= len(s); i++ {
		part := s[start:i]
		if flag[start][i-1] {
			path = append(path, part)
			backtraceV2(s, i, flag, path, ans)
			path = path[0 : len(path)-1]
		}
	}
}

func partitionV1(s string) [][]string {
	var path []string
	var ans [][]string
	backtraceV1(s, 0, path, &ans)
	return ans
}

func backtraceV1(s string, start int, path []string, ans *[][]string) {
	if start == len(s) {
		*ans = append(*ans, append([]string{}, path...))
		return
	}

	// 需要注意：这里需要包含i等于len(s)的情况
	for i := start + 1; i <= len(s); i++ {
		part := s[start:i]
		if check(part) {
			path = append(path, part)
			backtraceV1(s, i, path, ans)
			path = path[0 : len(path)-1]
		}
	}
}

func check(s string) bool {
	l, r := 0, len(s)-1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}
