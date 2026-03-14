package p0104_maximum_depth_of_binary_tree

import "testing"

func TestMaxDepth(t *testing.T) {
	root := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	if got := maxDepth(root); got != 3 {
		t.Errorf("maxDepth = %d, want 3", got)
	}
	if got := maxDepth(nil); got != 0 {
		t.Errorf("maxDepth(nil) = %d, want 0", got)
	}
}
