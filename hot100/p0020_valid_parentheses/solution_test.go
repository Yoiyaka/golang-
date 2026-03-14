package p0020_valid_parentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}
	for _, tc := range tests {
		got := isValid(tc.s)
		if got != tc.want {
			t.Errorf("isValid(%q) = %v, want %v", tc.s, got, tc.want)
		}
	}
}
