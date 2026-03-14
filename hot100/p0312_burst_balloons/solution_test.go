package p0312_burst_balloons

import "testing"

func TestMaxCoins(t *testing.T) {
	if got := maxCoins([]int{3, 1, 5, 8}); got != 167 {
		t.Errorf("got %d, want 167", got)
	}
	if got := maxCoins([]int{1, 5}); got != 10 {
		t.Errorf("got %d, want 10", got)
	}
}
