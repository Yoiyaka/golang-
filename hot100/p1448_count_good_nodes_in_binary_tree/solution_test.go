package p1448_count_good_nodes_in_binary_tree

import "testing"

func TestGoodNodes(t *testing.T) {
	root := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 5}},
	}
	if got := goodNodes(root); got != 4 {
		t.Errorf("got %d, want 4", got)
	}
	if got := goodNodes(&TreeNode{Val: 1}); got != 1 {
		t.Errorf("got %d, want 1", got)
	}
}
