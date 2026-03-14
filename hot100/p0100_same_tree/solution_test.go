package p0100_same_tree

import "testing"

func TestIsSameTree(t *testing.T) {
	p := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	q := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	if !isSameTree(p, q) {
		t.Error("expected same")
	}

	p2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	q2 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	if isSameTree(p2, q2) {
		t.Error("expected different")
	}
}
