package p0098_validate_binary_search_tree

import "testing"

func TestIsValidBST(t *testing.T) {
	valid := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	if !isValidBST(valid) {
		t.Error("expected valid BST")
	}

	invalid := &TreeNode{Val: 5,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}},
	}
	if isValidBST(invalid) {
		t.Error("expected invalid BST")
	}
}
