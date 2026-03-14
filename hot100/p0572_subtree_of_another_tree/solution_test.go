package p0572_subtree_of_another_tree

import "testing"

func TestIsSubtree(t *testing.T) {
	root := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
		Right: &TreeNode{Val: 5},
	}
	sub := &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}
	if !isSubtree(root, sub) {
		t.Error("expected subtree")
	}

	sub2 := &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}}}
	if isSubtree(root, sub2) {
		t.Error("expected not subtree")
	}
}
