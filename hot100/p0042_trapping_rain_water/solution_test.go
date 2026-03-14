package p0042_trapping_rain_water

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
		{[]int{3, 0, 2, 0, 4}, 7},
	}
	for _, tc := range tests {
		got := trap(tc.height)
		if got != tc.want {
			t.Errorf("trap(%v) = %d, want %d", tc.height, got, tc.want)
		}
	}
}
