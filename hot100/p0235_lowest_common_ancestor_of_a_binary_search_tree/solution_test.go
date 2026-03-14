package p0235_lowest_common_ancestor_of_a_binary_search_tree

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	n2 := &TreeNode{Val: 2}
	n8 := &TreeNode{Val: 8}
	root := &TreeNode{Val: 6,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}},
		Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}},
	}
	// find nodes by value
	var find func(*TreeNode, int) *TreeNode
	find = func(node *TreeNode, val int) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Val == val {
			return node
		}
		if l := find(node.Left, val); l != nil {
			return l
		}
		return find(node.Right, val)
	}
	n2 = find(root, 2)
	n8 = find(root, 8)
	got := lowestCommonAncestor(root, n2, n8)
	if got.Val != 6 {
		t.Errorf("got %d, want 6", got.Val)
	}
	p := find(root, 2)
	q := find(root, 4)
	got2 := lowestCommonAncestor(root, p, q)
	if got2.Val != 2 {
		t.Errorf("got %d, want 2", got2.Val)
	}
}
