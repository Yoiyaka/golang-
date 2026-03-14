package p0131_palindrome_partitioning

func partition(s string) [][]string {
	result := [][]string{}
	var backtrack func(start int, curr []string)
	backtrack = func(start int, curr []string) {
		if start == len(s) {
			tmp := make([]string, len(curr))
			copy(tmp, curr)
			result = append(result, tmp)
			return
		}
		for end := start + 1; end <= len(s); end++ {
			sub := s[start:end]
			if isPalin(sub) {
				backtrack(end, append(curr, sub))
			}
		}
	}
	backtrack(0, []string{})
	return result
}

func isPalin(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
