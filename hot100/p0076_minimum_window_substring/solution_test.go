package p0076_minimum_window_substring

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		s, tt, want string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
	}
	for _, tc := range tests {
		got := minWindow(tc.s, tc.tt)
		if got != tc.want {
			t.Errorf("minWindow(%q,%q) = %q, want %q", tc.s, tc.tt, got, tc.want)
		}
	}
}
