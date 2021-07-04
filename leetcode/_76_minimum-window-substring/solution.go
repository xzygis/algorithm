package _76_minimum_window_substring

func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	left, right := 0, 0
	resLen := len(s) + 1
	resLeft, resRight := -1, -1

	for left <= right && right < len(s) {
		if ori[s[right]] > 0 {
			cnt[s[right]]++
			for check(ori, cnt) && left <= right {
				newLen := right - left + 1
				if newLen < resLen {
					resLen = newLen
					resLeft = left
					resRight = right
				}

				if _, ok := cnt[s[left]]; ok {
					cnt[s[left]]--
				}
				left++
			}
		}
		right++
	}

	if resLeft == -1 {
		return ""
	}

	return s[resLeft:resRight+1]
}

func check(ori, cnt map[byte]int) bool {
	for k, v := range ori {
		if cnt[k] < v {
			return false
		}
	}
	return true
}
