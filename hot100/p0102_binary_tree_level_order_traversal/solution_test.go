package p0102_binary_tree_level_order_traversal

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	root := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	got := levelOrder(root)
	want := [][]int{{3}, {9, 20}, {15, 7}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
	if levelOrder(nil) != nil {
		t.Error("expected nil for nil input")
	}
}
