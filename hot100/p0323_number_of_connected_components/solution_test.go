package p0323_number_of_connected_components

import "testing"

func TestCountComponents(t *testing.T) {
	if got := countComponents(5, [][]int{{0, 1}, {1, 2}, {3, 4}}); got != 2 {
		t.Errorf("got %d, want 2", got)
	}
	if got := countComponents(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}); got != 1 {
		t.Errorf("got %d, want 1", got)
	}
}
