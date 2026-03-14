package p0230_kth_smallest_element_in_a_bst

import "testing"

func TestKthSmallest(t *testing.T) {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}}
	if got := kthSmallest(root, 1); got != 1 {
		t.Errorf("got %d, want 1", got)
	}
	root2 := &TreeNode{Val: 5,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 6},
	}
	if got := kthSmallest(root2, 3); got != 3 {
		t.Errorf("got %d, want 3", got)
	}
}
