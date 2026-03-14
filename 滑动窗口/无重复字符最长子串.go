package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int)
	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		ch := s[right]
		if idx, found := charIndex[ch]; found && idx >= left {
			left = idx + 1
		}
		charIndex[ch] = right
		if maxLen < right-left+1 {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 输出: 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 输出: 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 输出: 3
}
