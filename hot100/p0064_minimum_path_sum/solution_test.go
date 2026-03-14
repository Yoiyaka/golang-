package p0064_minimum_path_sum

import "testing"

func TestMinPathSum(t *testing.T) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	if got := minPathSum(grid); got != 7 {
		t.Errorf("got %d, want 7", got)
	}
	grid2 := [][]int{{1, 2, 3}, {4, 5, 6}}
	if got := minPathSum(grid2); got != 12 {
		t.Errorf("got %d, want 12", got)
	}
}
