package p0070_climbing_stairs

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct{ n, want int }{{2, 2}, {3, 3}, {5, 8}}
	for _, tc := range tests {
		if got := climbStairs(tc.n); got != tc.want {
			t.Errorf("climbStairs(%d) = %d, want %d", tc.n, got, tc.want)
		}
	}
}
