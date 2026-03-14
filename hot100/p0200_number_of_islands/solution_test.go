package p0200_number_of_islands

import "testing"

func TestNumIslands(t *testing.T) {
	grid1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	if got := numIslands(grid1); got != 1 {
		t.Errorf("got %d, want 1", got)
	}
	grid2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	if got := numIslands(grid2); got != 3 {
		t.Errorf("got %d, want 3", got)
	}
}
