package p0039_combination_sum

import "testing"

func TestCombinationSum(t *testing.T) {
	got := combinationSum([]int{2, 3, 6, 7}, 7)
	if len(got) != 2 {
		t.Errorf("expected 2 combinations, got %d: %v", len(got), got)
	}
	got2 := combinationSum([]int{2, 3, 5}, 8)
	if len(got2) != 3 {
		t.Errorf("expected 3 combinations, got %d: %v", len(got2), got2)
	}
}
