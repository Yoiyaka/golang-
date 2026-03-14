package p0094_binary_tree_inorder_traversal

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	got := inorderTraversal(root)
	want := []int{1, 3, 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
	if len(inorderTraversal(nil)) != 0 {
		t.Error("expected empty for nil")
	}
}
