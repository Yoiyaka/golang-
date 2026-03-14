package p0078_subsets

import "testing"

func TestSubsets(t *testing.T) {
	got := subsets([]int{1, 2, 3})
	if len(got) != 8 {
		t.Errorf("expected 8 subsets, got %d: %v", len(got), got)
	}
	got2 := subsets([]int{0})
	if len(got2) != 2 {
		t.Errorf("expected 2 subsets for [0], got %d", len(got2))
	}
}
