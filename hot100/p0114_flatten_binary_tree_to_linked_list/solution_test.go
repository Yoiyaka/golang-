package p0114_flatten_binary_tree_to_linked_list

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}},
	}
	flatten(root)
	result := []int{}
	curr := root
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Right
	}
	want := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("got %v, want %v", result, want)
	}
}
