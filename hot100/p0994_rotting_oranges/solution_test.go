package p0994_rotting_oranges

import "testing"

func TestOrangesRotting(t *testing.T) {
	grid1 := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	if got := orangesRotting(grid1); got != 4 {
		t.Errorf("got %d, want 4", got)
	}
	grid2 := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
	if got := orangesRotting(grid2); got != -1 {
		t.Errorf("got %d, want -1", got)
	}
	grid3 := [][]int{{0, 2}}
	if got := orangesRotting(grid3); got != 0 {
		t.Errorf("got %d, want 0", got)
	}
}
