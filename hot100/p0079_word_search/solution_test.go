package p0079_word_search

import "testing"

func TestExist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	if !exist(board, "ABCCED") {
		t.Error("expected true for ABCCED")
	}
	if !exist(board, "SEE") {
		t.Error("expected true for SEE")
	}
	if exist(board, "ABCB") {
		t.Error("expected false for ABCB")
	}
}
