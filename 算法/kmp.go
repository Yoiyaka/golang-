func kmp(text, pattern string) int {
	if len(pattern) == 0 {
		return 0
	}

	// 构建Next数组
	next := make([]int, len(pattern))
	j := 0
	for i := 1; i < len(pattern); i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = next[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		next[i] = j
	}

	// 匹配过程
	j = 0
	for i := 0; i < len(text); i++ {
		for j > 0 && text[i] != pattern[j] {
			j = next[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == len(pattern) {
			return i - j + 1
		}
	}
	return -1
}