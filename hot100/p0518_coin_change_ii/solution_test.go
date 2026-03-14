package p0518_coin_change_ii

import "testing"

func TestChange(t *testing.T) {
	if got := change(5, []int{1, 2, 5}); got != 4 {
		t.Errorf("got %d, want 4", got)
	}
	if got := change(3, []int{2}); got != 0 {
		t.Errorf("got %d, want 0", got)
	}
	if got := change(10, []int{10}); got != 1 {
		t.Errorf("got %d, want 1", got)
	}
}
