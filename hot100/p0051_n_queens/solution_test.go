package p0051_n_queens

import "testing"

func TestSolveNQueens(t *testing.T) {
	got := solveNQueens(4)
	if len(got) != 2 {
		t.Errorf("expected 2 solutions for n=4, got %d", len(got))
	}
	got1 := solveNQueens(1)
	if len(got1) != 1 {
		t.Errorf("expected 1 solution for n=1, got %d", len(got1))
	}
}
