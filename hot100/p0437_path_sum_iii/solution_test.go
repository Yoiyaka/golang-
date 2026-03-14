package p0437_path_sum_iii

import "testing"

func TestPathSum(t *testing.T) {
	root := &TreeNode{Val: 10,
		Left: &TreeNode{Val: 5,
			Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: -2}},
			Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 1}},
		},
		Right: &TreeNode{Val: -3, Right: &TreeNode{Val: 11}},
	}
	if got := pathSum(root, 8); got != 3 {
		t.Errorf("got %d, want 3", got)
	}
	if got := pathSum(nil, 0); got != 0 {
		t.Errorf("got %d, want 0", got)
	}
}
