package p0003_longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int)
	max := 0
	left := 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		if idx, ok := charIndex[c]; ok && idx >= left {
			left = idx + 1
		}
		charIndex[c] = right
		if right-left+1 > max {
			max = right - left + 1
		}
	}
	return max
}
