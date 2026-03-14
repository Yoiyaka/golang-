package p0124_binary_tree_maximum_path_sum

import "testing"

func TestMaxPathSum(t *testing.T) {
	root1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	if got := maxPathSum(root1); got != 6 {
		t.Errorf("got %d, want 6", got)
	}

	root2 := &TreeNode{Val: -10,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	if got := maxPathSum(root2); got != 42 {
		t.Errorf("got %d, want 42", got)
	}
}
