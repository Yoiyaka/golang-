package p0091_decode_ways

import "testing"

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"12", 2},
		{"226", 3},
		{"06", 0},
		{"11106", 2},
	}
	for _, tc := range tests {
		if got := numDecodings(tc.s); got != tc.want {
			t.Errorf("numDecodings(%q) = %d, want %d", tc.s, got, tc.want)
		}
	}
}
