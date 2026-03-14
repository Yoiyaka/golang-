package p1143_longest_common_subsequence

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		t1, t2 string
		want   int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
	}
	for _, tc := range tests {
		if got := longestCommonSubsequence(tc.t1, tc.t2); got != tc.want {
			t.Errorf("lcs(%q,%q) = %d, want %d", tc.t1, tc.t2, got, tc.want)
		}
	}
}
