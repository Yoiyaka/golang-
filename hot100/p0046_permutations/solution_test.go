package p0046_permutations

import "testing"

func TestPermute(t *testing.T) {
	got := permute([]int{1, 2, 3})
	if len(got) != 6 {
		t.Errorf("expected 6 permutations, got %d", len(got))
	}
	got2 := permute([]int{1})
	if len(got2) != 1 || got2[0][0] != 1 {
		t.Errorf("expected [[1]], got %v", got2)
	}
}
