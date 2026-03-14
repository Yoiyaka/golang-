package p0417_pacific_atlantic_water_flow

import "testing"

func TestPacificAtlantic(t *testing.T) {
	heights := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	got := pacificAtlantic(heights)
	if len(got) != 7 {
		t.Errorf("expected 7 cells, got %d: %v", len(got), got)
	}
}
