package p0101_symmetric_tree

import "testing"

func TestIsSymmetric(t *testing.T) {
	sym := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}},
	}
	if !isSymmetric(sym) {
		t.Error("expected symmetric")
	}

	asym := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
	}
	if isSymmetric(asym) {
		t.Error("expected asymmetric")
	}
}
