package p0003_longest_substring_without_repeating_characters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
	}
	for _, tc := range tests {
		got := lengthOfLongestSubstring(tc.s)
		if got != tc.want {
			t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", tc.s, got, tc.want)
		}
	}
}
