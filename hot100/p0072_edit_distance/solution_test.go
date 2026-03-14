package p0072_edit_distance

import "testing"

func TestMinDistance(t *testing.T) {
	tests := []struct {
		w1, w2 string
		want   int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
		{"", "", 0},
	}
	for _, tc := range tests {
		if got := minDistance(tc.w1, tc.w2); got != tc.want {
			t.Errorf("minDistance(%q,%q) = %d, want %d", tc.w1, tc.w2, got, tc.want)
		}
	}
}
