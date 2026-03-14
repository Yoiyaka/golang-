package p0543_diameter_of_binary_tree

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3},
	}
	if got := diameterOfBinaryTree(root); got != 3 {
		t.Errorf("got %d, want 3", got)
	}
	if got := diameterOfBinaryTree(&TreeNode{Val: 1}); got != 0 {
		t.Errorf("got %d, want 0", got)
	}
}
