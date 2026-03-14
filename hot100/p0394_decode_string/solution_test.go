package p0394_decode_string

import "testing"

func TestDecodeString(t *testing.T) {
	tests := []struct {
		s, want string
	}{
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
	}
	for _, tc := range tests {
		got := decodeString(tc.s)
		if got != tc.want {
			t.Errorf("decodeString(%q) = %q, want %q", tc.s, got, tc.want)
		}
	}
}
