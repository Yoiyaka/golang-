package p0010_regular_expression_matching

import "testing"

func TestIsMatch(t *testing.T) {
	tests := []struct {
		s, p string
		want bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
	}
	for _, tc := range tests {
		if got := isMatch(tc.s, tc.p); got != tc.want {
			t.Errorf("isMatch(%q,%q) = %v, want %v", tc.s, tc.p, got, tc.want)
		}
	}
}
