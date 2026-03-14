package p0062_unique_paths

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct{ m, n, want int }{{3, 7, 28}, {3, 2, 3}, {1, 1, 1}}
	for _, tc := range tests {
		if got := uniquePaths(tc.m, tc.n); got != tc.want {
			t.Errorf("uniquePaths(%d,%d) = %d, want %d", tc.m, tc.n, got, tc.want)
		}
	}
}
