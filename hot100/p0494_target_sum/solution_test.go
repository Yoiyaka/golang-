package p0494_target_sum

import "testing"

func TestFindTargetSumWays(t *testing.T) {
	if got := findTargetSumWays([]int{1, 1, 1, 1, 1}, 3); got != 5 {
		t.Errorf("got %d, want 5", got)
	}
	if got := findTargetSumWays([]int{1}, 1); got != 1 {
		t.Errorf("got %d, want 1", got)
	}
}
