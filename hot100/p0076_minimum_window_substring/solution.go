package p0076_minimum_window_substring

func minWindow(s string, t string) string {
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	have, total := 0, len(need)
	window := make(map[byte]int)
	res := ""
	resLen := len(s) + 1
	left := 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		window[c]++
		if need[c] > 0 && window[c] == need[c] {
			have++
		}
		for have == total {
			if right-left+1 < resLen {
				resLen = right - left + 1
				res = s[left : right+1]
			}
			lc := s[left]
			window[lc]--
			if need[lc] > 0 && window[lc] < need[lc] {
				have--
			}
			left++
		}
	}
	return res
}
