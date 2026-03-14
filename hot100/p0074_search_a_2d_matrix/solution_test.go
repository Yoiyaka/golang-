package p0074_search_a_2d_matrix

import "testing"

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	if !searchMatrix(matrix, 3) {
		t.Error("expected true for target 3")
	}
	if searchMatrix(matrix, 13) {
		t.Error("expected false for target 13")
	}
}
