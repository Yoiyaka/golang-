package p0221_maximal_square

import "testing"

func TestMaximalSquare(t *testing.T) {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	if got := maximalSquare(matrix); got != 4 {
		t.Errorf("got %d, want 4", got)
	}
	m2 := [][]byte{{'0', '1'}, {'1', '0'}}
	if got := maximalSquare(m2); got != 1 {
		t.Errorf("got %d, want 1", got)
	}
}
