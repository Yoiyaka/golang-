package p0110_balanced_binary_tree

import "testing"

func TestIsBalanced(t *testing.T) {
	balanced := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	if !isBalanced(balanced) {
		t.Error("expected balanced")
	}

	unbalanced := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}},
	}
	if isBalanced(unbalanced) {
		t.Error("expected unbalanced")
	}
}
