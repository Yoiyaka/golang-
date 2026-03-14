func minWindow(s string, t string) string {
	need := make(map[byte]int) // 统计 t 中每个字符出现次数
	for i := range t {
		need[t[i]]++
	}
	window := make(map[byte]int) // 统计窗口中的字符
	left, right := 0, 0
	valid := 0
	minLen := len(s) + 1
	start := 0

	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// 收缩窗口，直到窗口内的所有 t 的字符都满足数量要求
		for valid == len(need) {
			if right-left < minLen {
				minLen = right - left
				start = left
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if minLen > len(s) {
		return ""
	}
	return s[start : start+minLen]
}